package handlers

import (
	"errors"
	"fmt"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"math/big"
	"net/http"
	"time"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"
	"gitlab.com/tokend/nft-books/generator-svc/internal/signature"
)

func SignMint(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewSignMintRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch new sign mint request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Getting task's mintInfo
	task, err := helpers.DB(r).Tasks().GetById(request.TaskID)
	if err != nil {
		logger.WithError(err).Error("failed to get task")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if task == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	// Getting book's mintInfo
	book, err := helpers.Booker(r).GetBookById(task.BookId)
	if err != nil {
		logger.WithError(err).Error("failed to get a book")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if book == nil {
		logger.Warnf("Book with specified id %d was not found", task.BookId)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	// Getting price per token in dollars
	priceResponse, err := helpers.Pricer(r).GetPrice(request.Platform, request.TokenAddress)
	if err != nil {
		logger.WithError(err).Error("failed to get price")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// Forming signature mintInfo
	mintConfig := helpers.Minter(r)

	domainData := signature.EIP712DomainData{
		VerifyingAddress: book.Data.Attributes.ContractAddress,
		ContractName:     book.Data.Attributes.ContractName,
		ContractVersion:  book.Data.Attributes.ContractVersion,
		ChainID:          book.Data.Attributes.ChainId,
	}

	mintInfo := signature.MintInfo{
		TokenAddress: request.TokenAddress,
		TokenURI:     task.MetadataIpfsHash,
	}

	mintInfo.PricePerOneToken, err = helpers.ConvertPrice(priceResponse.Data.Attributes.Price, mintConfig.Precision)
	if err != nil {
		logger.WithError(err).Error("failed to convert price")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	mintInfo.EndTimestamp = time.Now().Add(mintConfig.Expiration).Unix()
	mintInfo.Discount = big.NewInt(0)

	//Getting promocode info
	var promocode *data.Promocode = nil

	if request.PromocodeID != 0 {
		promocode, err = helpers.DB(r).Promocodes().FilterById(request.PromocodeID).Get()
		if err != nil {
			logger.WithError(err).Error("failed to get promocode")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		//Validating promocode
		promocodeResponse, err := responses.NewValidatePromocodeResponse(*promocode)

		if err != nil {
			logger.WithError(err).Error("failed to get promocode response")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		if promocodeResponse.Data.Attributes.State != resources.PromocodeActive {
			logger.WithError(err).Error("promocode invalid")
			ape.RenderErr(w, problems.BadRequest(errors.New("promocode invalid"))...)
			return
		}

		//discount in contract is int number where 1% = 10^25
		//discount in data base is float number where 1% = 0.01
		const coefString string = "10000000000000000000000000"
		discount := big.NewInt(int64(promocode.Discount * 100))
		coef, ok := big.NewInt(0).SetString(fmt.Sprintf("%v", coefString), 10)
		if !ok {
			logger.Error("failed to parse big int")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		discount.Mul(discount, coef)

		mintInfo.Discount = discount
	}

	// Signing the mint transaction
	mintSignature, err := signature.SignMintInfo(&mintInfo, &domainData, &mintConfig)
	if err != nil {
		logger.WithError(err).Error("failed to generate eip712 mint signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewSignMintResponse(
		mintInfo.PricePerOneToken.String(),
		mintSignature,
		mintInfo.EndTimestamp,
	))
}
