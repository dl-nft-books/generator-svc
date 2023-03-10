//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package sampling ;import (_d "github.com/unidoc/unipdf/v3/internal/bitwise";_df "github.com/unidoc/unipdf/v3/internal/imageutil";_cg "io";);func (_e *Reader )ReadSample ()(uint32 ,error ){if _e ._db ==_e ._g .Height {return 0,_cg .EOF ;};_a ,_aa :=_e ._b .ReadBits (byte (_e ._g .BitsPerComponent ));
if _aa !=nil {return 0,_aa ;};_e ._dd --;if _e ._dd ==0{_e ._dd =_e ._g .ColorComponents ;_e ._ga ++;};if _e ._ga ==_e ._g .Width {if _e ._gb {_e ._b .ConsumeRemainingBits ();};_e ._ga =0;_e ._db ++;};return uint32 (_a ),nil ;};func (_ag *Writer )WriteSample (sample uint32 )error {if _ ,_cfb :=_ag ._aae .WriteBits (uint64 (sample ),_ag ._ed .BitsPerComponent );
_cfb !=nil {return _cfb ;};_ag ._fb --;if _ag ._fb ==0{_ag ._fb =_ag ._ed .ColorComponents ;_ag ._da ++;};if _ag ._da ==_ag ._ed .Width {if _ag ._bg {_ag ._aae .FinishByte ();};_ag ._da =0;};return nil ;};func NewWriter (img _df .ImageBase )*Writer {return &Writer {_aae :_d .NewWriterMSB (img .Data ),_ed :img ,_fb :img .ColorComponents ,_bg :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};type Reader struct{_g _df .ImageBase ;_b *_d .Reader ;_ga ,_db ,_dd int ;_gb bool ;};func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _cd []uint32 ;_cbb :=bitsPerSample ;var _fg uint32 ;var _ba byte ;_bf :=0;_ad :=0;_cgc :=0;for _cgc < len (data ){if _bf > 0{_bc :=_bf ;
if _cbb < _bc {_bc =_cbb ;};_fg =(_fg <<uint (_bc ))|uint32 (_ba >>uint (8-_bc ));_bf -=_bc ;if _bf > 0{_ba =_ba <<uint (_bc );}else {_ba =0;};_cbb -=_bc ;if _cbb ==0{_cd =append (_cd ,_fg );_cbb =bitsPerSample ;_fg =0;_ad ++;};}else {_dg :=data [_cgc ];
_cgc ++;_dc :=8;if _cbb < _dc {_dc =_cbb ;};_bf =8-_dc ;_fg =(_fg <<uint (_dc ))|uint32 (_dg >>uint (_bf ));if _dc < 8{_ba =_dg <<uint (_dc );};_cbb -=_dc ;if _cbb ==0{_cd =append (_cd ,_fg );_cbb =bitsPerSample ;_fg =0;_ad ++;};};};for _bf >=bitsPerSample {_de :=_bf ;
if _cbb < _de {_de =_cbb ;};_fg =(_fg <<uint (_de ))|uint32 (_ba >>uint (8-_de ));_bf -=_de ;if _bf > 0{_ba =_ba <<uint (_de );}else {_ba =0;};_cbb -=_de ;if _cbb ==0{_cd =append (_cd ,_fg );_cbb =bitsPerSample ;_fg =0;_ad ++;};};return _cd ;};type SampleReader interface{ReadSample ()(uint32 ,error );
ReadSamples (_f []uint32 )error ;};type SampleWriter interface{WriteSample (_bbe uint32 )error ;WriteSamples (_dca []uint32 )error ;};type Writer struct{_ed _df .ImageBase ;_aae *_d .Writer ;_da ,_fb int ;_bg bool ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _dff []uint32 ;
_fgf :=bitsPerOutputSample ;var _ddf uint32 ;var _bb uint32 ;_cf :=0;_fgc :=0;_aaf :=0;for _aaf < len (data ){if _cf > 0{_dee :=_cf ;if _fgf < _dee {_dee =_fgf ;};_ddf =(_ddf <<uint (_dee ))|(_bb >>uint (bitsPerInputSample -_dee ));_cf -=_dee ;if _cf > 0{_bb =_bb <<uint (_dee );
}else {_bb =0;};_fgf -=_dee ;if _fgf ==0{_dff =append (_dff ,_ddf );_fgf =bitsPerOutputSample ;_ddf =0;_fgc ++;};}else {_ac :=data [_aaf ];_aaf ++;_bad :=bitsPerInputSample ;if _fgf < _bad {_bad =_fgf ;};_cf =bitsPerInputSample -_bad ;_ddf =(_ddf <<uint (_bad ))|(_ac >>uint (_cf ));
if _bad < bitsPerInputSample {_bb =_ac <<uint (_bad );};_fgf -=_bad ;if _fgf ==0{_dff =append (_dff ,_ddf );_fgf =bitsPerOutputSample ;_ddf =0;_fgc ++;};};};for _cf >=bitsPerOutputSample {_dfe :=_cf ;if _fgf < _dfe {_dfe =_fgf ;};_ddf =(_ddf <<uint (_dfe ))|(_bb >>uint (bitsPerInputSample -_dfe ));
_cf -=_dfe ;if _cf > 0{_bb =_bb <<uint (_dfe );}else {_bb =0;};_fgf -=_dfe ;if _fgf ==0{_dff =append (_dff ,_ddf );_fgf =bitsPerOutputSample ;_ddf =0;_fgc ++;};};if _fgf > 0&&_fgf < bitsPerOutputSample {_ddf <<=uint (_fgf );_dff =append (_dff ,_ddf );};
return _dff ;};func NewReader (img _df .ImageBase )*Reader {return &Reader {_b :_d .NewReader (img .Data ),_g :img ,_dd :img .ColorComponents ,_gb :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};func (_age *Writer )WriteSamples (samples []uint32 )error {for _gd :=0;
_gd < len (samples );_gd ++{if _ec :=_age .WriteSample (samples [_gd ]);_ec !=nil {return _ec ;};};return nil ;};func (_cb *Reader )ReadSamples (samples []uint32 )(_af error ){for _cga :=0;_cga < len (samples );_cga ++{samples [_cga ],_af =_cb .ReadSample ();
if _af !=nil {return _af ;};};return nil ;};