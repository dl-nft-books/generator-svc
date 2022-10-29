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

package ccittfax ;import (_f "errors";_ff "github.com/unidoc/unipdf/v3/internal/bitwise";_c "io";_a "math";);func _bef (_babc int )([]byte ,int ){var _faf []byte ;for _aedd :=0;_aedd < 6;_aedd ++{_faf ,_babc =_adf (_faf ,_babc ,_cd );};return _faf ,_babc %8;
};func init (){_cf =&treeNode {_agd :true ,_gcbf :_ad };_ffb =&treeNode {_gcbf :_gc ,_aad :_cf };_ffb ._ggbe =_ffb ;_fg =&tree {_gadb :&treeNode {}};if _eg :=_fg .fillWithNode (12,0,_ffb );_eg !=nil {panic (_eg .Error ());};if _cef :=_fg .fillWithNode (12,1,_cf );
_cef !=nil {panic (_cef .Error ());};_ffa =&tree {_gadb :&treeNode {}};for _ee :=0;_ee < len (_gba );_ee ++{for _b :=0;_b < len (_gba [_ee ]);_b ++{if _fc :=_ffa .fill (_ee +2,int (_gba [_ee ][_b ]),int (_bb [_ee ][_b ]));_fc !=nil {panic (_fc .Error ());
};};};if _af :=_ffa .fillWithNode (12,0,_ffb );_af !=nil {panic (_af .Error ());};if _dg :=_ffa .fillWithNode (12,1,_cf );_dg !=nil {panic (_dg .Error ());};_e =&tree {_gadb :&treeNode {}};for _gcc :=0;_gcc < len (_afg );_gcc ++{for _ef :=0;_ef < len (_afg [_gcc ]);
_ef ++{if _ag :=_e .fill (_gcc +4,int (_afg [_gcc ][_ef ]),int (_faa [_gcc ][_ef ]));_ag !=nil {panic (_ag .Error ());};};};if _da :=_e .fillWithNode (12,0,_ffb );_da !=nil {panic (_da .Error ());};if _ab :=_e .fillWithNode (12,1,_cf );_ab !=nil {panic (_ab .Error ());
};_g =&tree {_gadb :&treeNode {}};if _ba :=_g .fill (4,1,_ce );_ba !=nil {panic (_ba .Error ());};if _dd :=_g .fill (3,1,_gb );_dd !=nil {panic (_dd .Error ());};if _fd :=_g .fill (1,1,0);_fd !=nil {panic (_fd .Error ());};if _aa :=_g .fill (3,3,1);_aa !=nil {panic (_aa .Error ());
};if _fa :=_g .fill (6,3,2);_fa !=nil {panic (_fa .Error ());};if _bc :=_g .fill (7,3,3);_bc !=nil {panic (_bc .Error ());};if _eb :=_g .fill (3,2,-1);_eb !=nil {panic (_eb .Error ());};if _ec :=_g .fill (6,2,-2);_ec !=nil {panic (_ec .Error ());};if _bcd :=_g .fill (7,2,-3);
_bcd !=nil {panic (_bcd .Error ());};};func (_dba tiffType )String ()string {switch _dba {case _eed :return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u004d\u006f\u0064i\u0066\u0069\u0065\u0064\u0048\u0075\u0066\u0066\u006d\u0061n\u0052\u006c\u0065";
case _fca :return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0034";case _eee :return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0036";default:return "\u0075n\u0064\u0065\u0066\u0069\u006e\u0065d";};};func _daa (_aabd []byte ,_faca bool ,_ccce int )(int ,int ){_eaea :=0;
for _ccce < len (_aabd ){if _faca {if _aabd [_ccce ]!=_dcb {break ;};}else {if _aabd [_ccce ]!=_cdbf {break ;};};_eaea ++;_ccce ++;};return _eaea ,_ccce ;};func (_dge *Decoder )looseFetchEOL ()(bool ,error ){_egg ,_gcgf :=_dge ._gab .ReadBits (12);if _gcgf !=nil {return false ,_gcgf ;
};switch _egg {case 0x1:return true ,nil ;case 0x0:for {_dbbe ,_acc :=_dge ._gab .ReadBool ();if _acc !=nil {return false ,_acc ;};if _dbbe {return true ,nil ;};};default:return false ,nil ;};};func (_cgd *Encoder )encodeG32D (_fgge [][]byte )[]byte {var _bdc []byte ;
var _gbc int ;for _cgda :=0;_cgda < len (_fgge );_cgda +=_cgd .K {if _cgd .Rows > 0&&!_cgd .EndOfBlock &&_cgda ==_cgd .Rows {break ;};_beeg ,_cag :=_fcd (_fgge [_cgda ],_gbc ,_cd );_bdc =_cgd .appendEncodedRow (_bdc ,_beeg ,_gbc );if _cgd .EncodedByteAlign {_cag =0;
};_gbc =_cag ;for _gcbe :=_cgda +1;_gcbe < (_cgda +_cgd .K )&&_gcbe < len (_fgge );_gcbe ++{if _cgd .Rows > 0&&!_cgd .EndOfBlock &&_gcbe ==_cgd .Rows {break ;};_dfg ,_aeda :=_adf (nil ,_gbc ,_ddg );var _gdf ,_dbc ,_cgdg int ;_baef :=-1;for _baef < len (_fgge [_gcbe ]){_gdf =_ece (_fgge [_gcbe ],_baef );
_dbc =_bbg (_fgge [_gcbe ],_fgge [_gcbe -1],_baef );_cgdg =_ece (_fgge [_gcbe -1],_dbc );if _cgdg < _gdf {_dfg ,_aeda =_dac (_dfg ,_aeda );_baef =_cgdg ;}else {if _a .Abs (float64 (_dbc -_gdf ))> 3{_dfg ,_aeda ,_baef =_fefc (_fgge [_gcbe ],_dfg ,_aeda ,_baef ,_gdf );
}else {_dfg ,_aeda =_gea (_dfg ,_aeda ,_gdf ,_dbc );_baef =_gdf ;};};};_bdc =_cgd .appendEncodedRow (_bdc ,_dfg ,_gbc );if _cgd .EncodedByteAlign {_aeda =0;};_gbc =_aeda %8;};};if _cgd .EndOfBlock {_cbe ,_ :=_bef (_gbc );_bdc =_cgd .appendEncodedRow (_bdc ,_cbe ,_gbc );
};return _bdc ;};func (_gcf *treeNode )set (_dbacb bool ,_fcde *treeNode ){if !_dbacb {_gcf ._ggbe =_fcde ;}else {_gcf ._aad =_fcde ;};};var _gba =[...][]uint16 {{0x2,0x3},{0x2,0x3},{0x2,0x3},{0x3},{0x4,0x5},{0x4,0x5,0x7},{0x4,0x7},{0x18},{0x17,0x18,0x37,0x8,0xf},{0x17,0x18,0x28,0x37,0x67,0x68,0x6c,0x8,0xc,0xd},{0x12,0x13,0x14,0x15,0x16,0x17,0x1c,0x1d,0x1e,0x1f,0x24,0x27,0x28,0x2b,0x2c,0x33,0x34,0x35,0x37,0x38,0x52,0x53,0x54,0x55,0x56,0x57,0x58,0x59,0x5a,0x5b,0x64,0x65,0x66,0x67,0x68,0x69,0x6a,0x6b,0x6c,0x6d,0xc8,0xc9,0xca,0xcb,0xcc,0xcd,0xd2,0xd3,0xd4,0xd5,0xd6,0xd7,0xda,0xdb},{0x4a,0x4b,0x4c,0x4d,0x52,0x53,0x54,0x55,0x5a,0x5b,0x64,0x65,0x6c,0x6d,0x72,0x73,0x74,0x75,0x76,0x77}};
func init (){_eeb =make (map[int ]code );_eeb [0]=code {Code :13<<8|3<<6,BitsWritten :10};_eeb [1]=code {Code :2<<(5+8),BitsWritten :3};_eeb [2]=code {Code :3<<(6+8),BitsWritten :2};_eeb [3]=code {Code :2<<(6+8),BitsWritten :2};_eeb [4]=code {Code :3<<(5+8),BitsWritten :3};
_eeb [5]=code {Code :3<<(4+8),BitsWritten :4};_eeb [6]=code {Code :2<<(4+8),BitsWritten :4};_eeb [7]=code {Code :3<<(3+8),BitsWritten :5};_eeb [8]=code {Code :5<<(2+8),BitsWritten :6};_eeb [9]=code {Code :4<<(2+8),BitsWritten :6};_eeb [10]=code {Code :4<<(1+8),BitsWritten :7};
_eeb [11]=code {Code :5<<(1+8),BitsWritten :7};_eeb [12]=code {Code :7<<(1+8),BitsWritten :7};_eeb [13]=code {Code :4<<8,BitsWritten :8};_eeb [14]=code {Code :7<<8,BitsWritten :8};_eeb [15]=code {Code :12<<8,BitsWritten :9};_eeb [16]=code {Code :5<<8|3<<6,BitsWritten :10};
_eeb [17]=code {Code :6<<8,BitsWritten :10};_eeb [18]=code {Code :2<<8,BitsWritten :10};_eeb [19]=code {Code :12<<8|7<<5,BitsWritten :11};_eeb [20]=code {Code :13<<8,BitsWritten :11};_eeb [21]=code {Code :13<<8|4<<5,BitsWritten :11};_eeb [22]=code {Code :6<<8|7<<5,BitsWritten :11};
_eeb [23]=code {Code :5<<8,BitsWritten :11};_eeb [24]=code {Code :2<<8|7<<5,BitsWritten :11};_eeb [25]=code {Code :3<<8,BitsWritten :11};_eeb [26]=code {Code :12<<8|10<<4,BitsWritten :12};_eeb [27]=code {Code :12<<8|11<<4,BitsWritten :12};_eeb [28]=code {Code :12<<8|12<<4,BitsWritten :12};
_eeb [29]=code {Code :12<<8|13<<4,BitsWritten :12};_eeb [30]=code {Code :6<<8|8<<4,BitsWritten :12};_eeb [31]=code {Code :6<<8|9<<4,BitsWritten :12};_eeb [32]=code {Code :6<<8|10<<4,BitsWritten :12};_eeb [33]=code {Code :6<<8|11<<4,BitsWritten :12};_eeb [34]=code {Code :13<<8|2<<4,BitsWritten :12};
_eeb [35]=code {Code :13<<8|3<<4,BitsWritten :12};_eeb [36]=code {Code :13<<8|4<<4,BitsWritten :12};_eeb [37]=code {Code :13<<8|5<<4,BitsWritten :12};_eeb [38]=code {Code :13<<8|6<<4,BitsWritten :12};_eeb [39]=code {Code :13<<8|7<<4,BitsWritten :12};_eeb [40]=code {Code :6<<8|12<<4,BitsWritten :12};
_eeb [41]=code {Code :6<<8|13<<4,BitsWritten :12};_eeb [42]=code {Code :13<<8|10<<4,BitsWritten :12};_eeb [43]=code {Code :13<<8|11<<4,BitsWritten :12};_eeb [44]=code {Code :5<<8|4<<4,BitsWritten :12};_eeb [45]=code {Code :5<<8|5<<4,BitsWritten :12};_eeb [46]=code {Code :5<<8|6<<4,BitsWritten :12};
_eeb [47]=code {Code :5<<8|7<<4,BitsWritten :12};_eeb [48]=code {Code :6<<8|4<<4,BitsWritten :12};_eeb [49]=code {Code :6<<8|5<<4,BitsWritten :12};_eeb [50]=code {Code :5<<8|2<<4,BitsWritten :12};_eeb [51]=code {Code :5<<8|3<<4,BitsWritten :12};_eeb [52]=code {Code :2<<8|4<<4,BitsWritten :12};
_eeb [53]=code {Code :3<<8|7<<4,BitsWritten :12};_eeb [54]=code {Code :3<<8|8<<4,BitsWritten :12};_eeb [55]=code {Code :2<<8|7<<4,BitsWritten :12};_eeb [56]=code {Code :2<<8|8<<4,BitsWritten :12};_eeb [57]=code {Code :5<<8|8<<4,BitsWritten :12};_eeb [58]=code {Code :5<<8|9<<4,BitsWritten :12};
_eeb [59]=code {Code :2<<8|11<<4,BitsWritten :12};_eeb [60]=code {Code :2<<8|12<<4,BitsWritten :12};_eeb [61]=code {Code :5<<8|10<<4,BitsWritten :12};_eeb [62]=code {Code :6<<8|6<<4,BitsWritten :12};_eeb [63]=code {Code :6<<8|7<<4,BitsWritten :12};_cefc =make (map[int ]code );
_cefc [0]=code {Code :53<<8,BitsWritten :8};_cefc [1]=code {Code :7<<(2+8),BitsWritten :6};_cefc [2]=code {Code :7<<(4+8),BitsWritten :4};_cefc [3]=code {Code :8<<(4+8),BitsWritten :4};_cefc [4]=code {Code :11<<(4+8),BitsWritten :4};_cefc [5]=code {Code :12<<(4+8),BitsWritten :4};
_cefc [6]=code {Code :14<<(4+8),BitsWritten :4};_cefc [7]=code {Code :15<<(4+8),BitsWritten :4};_cefc [8]=code {Code :19<<(3+8),BitsWritten :5};_cefc [9]=code {Code :20<<(3+8),BitsWritten :5};_cefc [10]=code {Code :7<<(3+8),BitsWritten :5};_cefc [11]=code {Code :8<<(3+8),BitsWritten :5};
_cefc [12]=code {Code :8<<(2+8),BitsWritten :6};_cefc [13]=code {Code :3<<(2+8),BitsWritten :6};_cefc [14]=code {Code :52<<(2+8),BitsWritten :6};_cefc [15]=code {Code :53<<(2+8),BitsWritten :6};_cefc [16]=code {Code :42<<(2+8),BitsWritten :6};_cefc [17]=code {Code :43<<(2+8),BitsWritten :6};
_cefc [18]=code {Code :39<<(1+8),BitsWritten :7};_cefc [19]=code {Code :12<<(1+8),BitsWritten :7};_cefc [20]=code {Code :8<<(1+8),BitsWritten :7};_cefc [21]=code {Code :23<<(1+8),BitsWritten :7};_cefc [22]=code {Code :3<<(1+8),BitsWritten :7};_cefc [23]=code {Code :4<<(1+8),BitsWritten :7};
_cefc [24]=code {Code :40<<(1+8),BitsWritten :7};_cefc [25]=code {Code :43<<(1+8),BitsWritten :7};_cefc [26]=code {Code :19<<(1+8),BitsWritten :7};_cefc [27]=code {Code :36<<(1+8),BitsWritten :7};_cefc [28]=code {Code :24<<(1+8),BitsWritten :7};_cefc [29]=code {Code :2<<8,BitsWritten :8};
_cefc [30]=code {Code :3<<8,BitsWritten :8};_cefc [31]=code {Code :26<<8,BitsWritten :8};_cefc [32]=code {Code :27<<8,BitsWritten :8};_cefc [33]=code {Code :18<<8,BitsWritten :8};_cefc [34]=code {Code :19<<8,BitsWritten :8};_cefc [35]=code {Code :20<<8,BitsWritten :8};
_cefc [36]=code {Code :21<<8,BitsWritten :8};_cefc [37]=code {Code :22<<8,BitsWritten :8};_cefc [38]=code {Code :23<<8,BitsWritten :8};_cefc [39]=code {Code :40<<8,BitsWritten :8};_cefc [40]=code {Code :41<<8,BitsWritten :8};_cefc [41]=code {Code :42<<8,BitsWritten :8};
_cefc [42]=code {Code :43<<8,BitsWritten :8};_cefc [43]=code {Code :44<<8,BitsWritten :8};_cefc [44]=code {Code :45<<8,BitsWritten :8};_cefc [45]=code {Code :4<<8,BitsWritten :8};_cefc [46]=code {Code :5<<8,BitsWritten :8};_cefc [47]=code {Code :10<<8,BitsWritten :8};
_cefc [48]=code {Code :11<<8,BitsWritten :8};_cefc [49]=code {Code :82<<8,BitsWritten :8};_cefc [50]=code {Code :83<<8,BitsWritten :8};_cefc [51]=code {Code :84<<8,BitsWritten :8};_cefc [52]=code {Code :85<<8,BitsWritten :8};_cefc [53]=code {Code :36<<8,BitsWritten :8};
_cefc [54]=code {Code :37<<8,BitsWritten :8};_cefc [55]=code {Code :88<<8,BitsWritten :8};_cefc [56]=code {Code :89<<8,BitsWritten :8};_cefc [57]=code {Code :90<<8,BitsWritten :8};_cefc [58]=code {Code :91<<8,BitsWritten :8};_cefc [59]=code {Code :74<<8,BitsWritten :8};
_cefc [60]=code {Code :75<<8,BitsWritten :8};_cefc [61]=code {Code :50<<8,BitsWritten :8};_cefc [62]=code {Code :51<<8,BitsWritten :8};_cefc [63]=code {Code :52<<8,BitsWritten :8};_eeg =make (map[int ]code );_eeg [64]=code {Code :3<<8|3<<6,BitsWritten :10};
_eeg [128]=code {Code :12<<8|8<<4,BitsWritten :12};_eeg [192]=code {Code :12<<8|9<<4,BitsWritten :12};_eeg [256]=code {Code :5<<8|11<<4,BitsWritten :12};_eeg [320]=code {Code :3<<8|3<<4,BitsWritten :12};_eeg [384]=code {Code :3<<8|4<<4,BitsWritten :12};
_eeg [448]=code {Code :3<<8|5<<4,BitsWritten :12};_eeg [512]=code {Code :3<<8|12<<3,BitsWritten :13};_eeg [576]=code {Code :3<<8|13<<3,BitsWritten :13};_eeg [640]=code {Code :2<<8|10<<3,BitsWritten :13};_eeg [704]=code {Code :2<<8|11<<3,BitsWritten :13};
_eeg [768]=code {Code :2<<8|12<<3,BitsWritten :13};_eeg [832]=code {Code :2<<8|13<<3,BitsWritten :13};_eeg [896]=code {Code :3<<8|18<<3,BitsWritten :13};_eeg [960]=code {Code :3<<8|19<<3,BitsWritten :13};_eeg [1024]=code {Code :3<<8|20<<3,BitsWritten :13};
_eeg [1088]=code {Code :3<<8|21<<3,BitsWritten :13};_eeg [1152]=code {Code :3<<8|22<<3,BitsWritten :13};_eeg [1216]=code {Code :119<<3,BitsWritten :13};_eeg [1280]=code {Code :2<<8|18<<3,BitsWritten :13};_eeg [1344]=code {Code :2<<8|19<<3,BitsWritten :13};
_eeg [1408]=code {Code :2<<8|20<<3,BitsWritten :13};_eeg [1472]=code {Code :2<<8|21<<3,BitsWritten :13};_eeg [1536]=code {Code :2<<8|26<<3,BitsWritten :13};_eeg [1600]=code {Code :2<<8|27<<3,BitsWritten :13};_eeg [1664]=code {Code :3<<8|4<<3,BitsWritten :13};
_eeg [1728]=code {Code :3<<8|5<<3,BitsWritten :13};_db =make (map[int ]code );_db [64]=code {Code :27<<(3+8),BitsWritten :5};_db [128]=code {Code :18<<(3+8),BitsWritten :5};_db [192]=code {Code :23<<(2+8),BitsWritten :6};_db [256]=code {Code :55<<(1+8),BitsWritten :7};
_db [320]=code {Code :54<<8,BitsWritten :8};_db [384]=code {Code :55<<8,BitsWritten :8};_db [448]=code {Code :100<<8,BitsWritten :8};_db [512]=code {Code :101<<8,BitsWritten :8};_db [576]=code {Code :104<<8,BitsWritten :8};_db [640]=code {Code :103<<8,BitsWritten :8};
_db [704]=code {Code :102<<8,BitsWritten :9};_db [768]=code {Code :102<<8|1<<7,BitsWritten :9};_db [832]=code {Code :105<<8,BitsWritten :9};_db [896]=code {Code :105<<8|1<<7,BitsWritten :9};_db [960]=code {Code :106<<8,BitsWritten :9};_db [1024]=code {Code :106<<8|1<<7,BitsWritten :9};
_db [1088]=code {Code :107<<8,BitsWritten :9};_db [1152]=code {Code :107<<8|1<<7,BitsWritten :9};_db [1216]=code {Code :108<<8,BitsWritten :9};_db [1280]=code {Code :108<<8|1<<7,BitsWritten :9};_db [1344]=code {Code :109<<8,BitsWritten :9};_db [1408]=code {Code :109<<8|1<<7,BitsWritten :9};
_db [1472]=code {Code :76<<8,BitsWritten :9};_db [1536]=code {Code :76<<8|1<<7,BitsWritten :9};_db [1600]=code {Code :77<<8,BitsWritten :9};_db [1664]=code {Code :24<<(2+8),BitsWritten :6};_db [1728]=code {Code :77<<8|1<<7,BitsWritten :9};_egf =make (map[int ]code );
_egf [1792]=code {Code :1<<8,BitsWritten :11};_egf [1856]=code {Code :1<<8|4<<5,BitsWritten :11};_egf [1920]=code {Code :1<<8|5<<5,BitsWritten :11};_egf [1984]=code {Code :1<<8|2<<4,BitsWritten :12};_egf [2048]=code {Code :1<<8|3<<4,BitsWritten :12};_egf [2112]=code {Code :1<<8|4<<4,BitsWritten :12};
_egf [2176]=code {Code :1<<8|5<<4,BitsWritten :12};_egf [2240]=code {Code :1<<8|6<<4,BitsWritten :12};_egf [2304]=code {Code :1<<8|7<<4,BitsWritten :12};_egf [2368]=code {Code :1<<8|12<<4,BitsWritten :12};_egf [2432]=code {Code :1<<8|13<<4,BitsWritten :12};
_egf [2496]=code {Code :1<<8|14<<4,BitsWritten :12};_egf [2560]=code {Code :1<<8|15<<4,BitsWritten :12};_agb =make (map[int ]byte );_agb [0]=0xFF;_agb [1]=0xFE;_agb [2]=0xFC;_agb [3]=0xF8;_agb [4]=0xF0;_agb [5]=0xE0;_agb [6]=0xC0;_agb [7]=0x80;_agb [8]=0x00;
};const (_ tiffType =iota ;_eed ;_fca ;_eee ;);type tiffType int ;func NewDecoder (data []byte ,options DecodeOptions )(*Decoder ,error ){_fed :=&Decoder {_gab :_ff .NewReader (data ),_ddgd :options .Columns ,_aea :options .Rows ,_cea :options .DamagedRowsBeforeError ,_cge :make ([]byte ,(options .Columns +7)/8),_cged :make ([]int ,options .Columns +2),_cc :make ([]int ,options .Columns +2),_ebf :options .EncodedByteAligned ,_cda :options .BlackIsOne ,_gbf :options .EndOfLine ,_ffba :options .EndOfBlock };
switch {case options .K ==0:_fed ._ca =_fca ;if len (data )< 20{return nil ,_f .New ("\u0074o\u006f\u0020\u0073\u0068o\u0072\u0074\u0020\u0063\u0063i\u0074t\u0066a\u0078\u0020\u0073\u0074\u0072\u0065\u0061m");};_gae :=data [:20];if _gae [0]!=0||(_gae [1]>>4!=1&&_gae [1]!=1){_fed ._ca =_eed ;
_dad :=(uint16 (_gae [0])<<8+uint16 (_gae [1]&0xff))>>4;for _dbd :=12;_dbd < 160;_dbd ++{_dad =(_dad <<1)+uint16 ((_gae [_dbd /8]>>uint16 (7-(_dbd %8)))&0x01);if _dad &0xfff==1{_fed ._ca =_fca ;break ;};};};case options .K < 0:_fed ._ca =_eee ;case options .K > 0:_fed ._ca =_fca ;
_fed ._dc =true ;};switch _fed ._ca {case _eed ,_fca ,_eee :default:return nil ,_f .New ("\u0075\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u0063\u0069\u0074\u0074\u0066\u0061\u0078\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0072\u0020ty\u0070\u0065");
};return _fed ,nil ;};func (_eggc *Encoder )encodeG31D (_dbac [][]byte )[]byte {var _gdb []byte ;_dgb :=0;for _cacd :=range _dbac {if _eggc .Rows > 0&&!_eggc .EndOfBlock &&_cacd ==_eggc .Rows {break ;};_fdg ,_aaba :=_fcd (_dbac [_cacd ],_dgb ,_fgg );_gdb =_eggc .appendEncodedRow (_gdb ,_fdg ,_dgb );
if _eggc .EncodedByteAlign {_aaba =0;};_dgb =_aaba ;};if _eggc .EndOfBlock {_caeg ,_ :=_ceaf (_dgb );_gdb =_eggc .appendEncodedRow (_gdb ,_caeg ,_dgb );};return _gdb ;};var (_ae =_f .New ("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0063\u006f\u0072\u0072\u0075\u0070\u0074\u0065\u0064\u0020\u0052T\u0043");
_bg =_f .New ("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0045\u004f\u004c\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064"););func _fcd (_gfg []byte ,_gaec int ,_ggf code )([]byte ,int ){_aedg :=true ;var _efa []byte ;_efa ,_gaec =_adf (nil ,_gaec ,_ggf );
_cce :=0;var _bba int ;for _cce < len (_gfg ){_bba ,_cce =_daa (_gfg ,_aedg ,_cce );_efa ,_gaec =_fff (_efa ,_gaec ,_bba ,_aedg );_aedg =!_aedg ;};return _efa ,_gaec %8;};func _ceaf (_gcd int )([]byte ,int ){var _egad []byte ;for _bf :=0;_bf < 6;_bf ++{_egad ,_gcd =_adf (_egad ,_gcd ,_fgg );
};return _egad ,_gcd %8;};func _ece (_acaa []byte ,_eef int )int {if _eef >=len (_acaa ){return _eef ;};if _eef < -1{_eef =-1;};var _dcc byte ;if _eef > -1{_dcc =_acaa [_eef ];}else {_dcc =_dcb ;};_dfcb :=_eef +1;for _dfcb < len (_acaa ){if _acaa [_dfcb ]!=_dcc {break ;
};_dfcb ++;};return _dfcb ;};type Encoder struct{K int ;EndOfLine bool ;EncodedByteAlign bool ;Columns int ;Rows int ;EndOfBlock bool ;BlackIs1 bool ;DamagedRowsBeforeError int ;};func (_ead *tree )fillWithNode (_dbge ,_cad int ,_eea *treeNode )error {_fcg :=_ead ._gadb ;
for _ddff :=0;_ddff < _dbge ;_ddff ++{_abeb :=uint (_dbge -1-_ddff );_caf :=((_cad >>_abeb )&1)!=0;_gbcb :=_fcg .walk (_caf );if _gbcb !=nil {if _gbcb ._agd {return _f .New ("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067");
};_fcg =_gbcb ;continue ;};if _ddff ==_dbge -1{_gbcb =_eea ;}else {_gbcb =&treeNode {};};if _cad ==0{_gbcb ._cgeff =true ;};_fcg .set (_caf ,_gbcb );_fcg =_gbcb ;};return nil ;};type Decoder struct{_ddgd int ;_aea int ;_abc int ;_cge []byte ;_cea int ;
_dc bool ;_beb bool ;_df bool ;_cda bool ;_gbf bool ;_ffba bool ;_ebf bool ;_cgc int ;_gg int ;_cged []int ;_cc []int ;_gga int ;_feb int ;_aae int ;_ge int ;_gab *_ff .Reader ;_ca tiffType ;_dfd error ;};func (_cgef *Decoder )fetch ()error {if _cgef ._cgc ==-1{return nil ;
};if _cgef ._gg < _cgef ._cgc {return nil ;};_cgef ._cgc =0;_cfg :=_cgef .decodeRow ();if _cfg !=nil {if !_f .Is (_cfg ,_c .EOF ){return _cfg ;};if _cgef ._cgc !=0{return _cfg ;};_cgef ._cgc =-1;};_cgef ._gg =0;return nil ;};func _bbg (_bgf ,_dcf []byte ,_gad int )int {_gec :=_ece (_dcf ,_gad );
if _gec < len (_dcf )&&(_gad ==-1&&_dcf [_gec ]==_dcb ||_gad >=0&&_gad < len (_bgf )&&_bgf [_gad ]==_dcf [_gec ]||_gad >=len (_bgf )&&_bgf [_gad -1]!=_dcf [_gec ]){_gec =_ece (_dcf ,_gec );};return _gec ;};func _acb (_ccg ,_gadg int )code {var _babd code ;
switch _gadg -_ccg {case -1:_babd =_gf ;case -2:_babd =_be ;case -3:_babd =_cg ;case 0:_babd =_cec ;case 1:_babd =_ed ;case 2:_babd =_bab ;case 3:_babd =_ga ;};return _babd ;};var _bb =[...][]uint16 {{3,2},{1,4},{6,5},{7},{9,8},{10,11,12},{13,14},{15},{16,17,0,18,64},{24,25,23,22,19,20,21,1792,1856,1920},{1984,2048,2112,2176,2240,2304,2368,2432,2496,2560,52,55,56,59,60,320,384,448,53,54,50,51,44,45,46,47,57,58,61,256,48,49,62,63,30,31,32,33,40,41,128,192,26,27,28,29,34,35,36,37,38,39,42,43},{640,704,768,832,1280,1344,1408,1472,1536,1600,1664,1728,512,576,896,960,1024,1088,1152,1216}};
func (_dag *Decoder )tryFetchRTC2D ()(_cgea error ){_dag ._gab .Mark ();var _adg bool ;for _agf :=0;_agf < 5;_agf ++{_adg ,_cgea =_dag .tryFetchEOL1 ();if _cgea !=nil {if _f .Is (_cgea ,_c .EOF ){if _agf ==0{break ;};return _ae ;};};if _adg {continue ;
};if _agf > 0{return _ae ;};break ;};if _adg {return _c .EOF ;};_dag ._gab .Reset ();return _cgea ;};var _faa =[...][]uint16 {{2,3,4,5,6,7},{128,8,9,64,10,11},{192,1664,16,17,13,14,15,1,12},{26,21,28,27,18,24,25,22,256,23,20,19},{33,34,35,36,37,38,31,32,29,53,54,39,40,41,42,43,44,30,61,62,63,0,320,384,45,59,60,46,49,50,51,52,55,56,57,58,448,512,640,576,47,48},{1472,1536,1600,1728,704,768,832,896,960,1024,1088,1152,1216,1280,1344,1408},{},{1792,1856,1920},{1984,2048,2112,2176,2240,2304,2368,2432,2496,2560}};
func _gabe (_gce int ,_dbdb bool )(code ,int ,bool ){if _gce < 64{if _dbdb {return _cefc [_gce ],0,true ;};return _eeb [_gce ],0,true ;};_ebgb :=_gce /64;if _ebgb > 40{return _egf [2560],_gce -2560,false ;};if _ebgb > 27{return _egf [_ebgb *64],_gce -_ebgb *64,false ;
};if _dbdb {return _db [_ebgb *64],_gce -_ebgb *64,false ;};return _eeg [_ebgb *64],_gce -_ebgb *64,false ;};func _adgd (_ecfd int )([]byte ,int ){var _gaa []byte ;for _gfa :=0;_gfa < 2;_gfa ++{_gaa ,_ecfd =_adf (_gaa ,_ecfd ,_fgg );};return _gaa ,_ecfd %8;
};func (_dfc *Decoder )decode1D ()error {var (_dfdb int ;_bee error ;);_bae :=true ;_dfc ._feb =0;for {var _dgg int ;if _bae {_dgg ,_bee =_dfc .decodeRun (_e );}else {_dgg ,_bee =_dfc .decodeRun (_ffa );};if _bee !=nil {return _bee ;};_dfdb +=_dgg ;_dfc ._cc [_dfc ._feb ]=_dfdb ;
_dfc ._feb ++;_bae =!_bae ;if _dfdb >=_dfc ._ddgd {break ;};};return nil ;};func (_bgabd *tree )fill (_ebgg ,_bgbb ,_eefd int )error {_gca :=_bgabd ._gadb ;for _acd :=0;_acd < _ebgg ;_acd ++{_ceba :=_ebgg -1-_acd ;_gddb :=((_bgbb >>uint (_ceba ))&1)!=0;
_caa :=_gca .walk (_gddb );if _caa !=nil {if _caa ._agd {return _f .New ("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067");
};_gca =_caa ;continue ;};_caa =&treeNode {};if _acd ==_ebgg -1{_caa ._gcbf =_eefd ;_caa ._agd =true ;};if _bgbb ==0{_caa ._cgeff =true ;};_gca .set (_gddb ,_caa );_gca =_caa ;};return nil ;};func (_gbab *Encoder )appendEncodedRow (_aecc ,_cecc []byte ,_fec int )[]byte {if len (_aecc )> 0&&_fec !=0&&!_gbab .EncodedByteAlign {_aecc [len (_aecc )-1]=_aecc [len (_aecc )-1]|_cecc [0];
_aecc =append (_aecc ,_cecc [1:]...);}else {_aecc =append (_aecc ,_cecc ...);};return _aecc ;};func (_eegb *Decoder )tryFetchEOL1 ()(bool ,error ){_gbg ,_fb :=_eegb ._gab .ReadBits (13);if _fb !=nil {return false ,_fb ;};return _gbg ==0x3,nil ;};func (_eae *Decoder )decode2D ()error {_eae ._gga =_eae ._feb ;
_eae ._cc ,_eae ._cged =_eae ._cged ,_eae ._cc ;_dce :=true ;var (_ccc bool ;_bcf int ;_gebc error ;);_eae ._feb =0;_cggg :for _bcf < _eae ._ddgd {_dfda :=_g ._gadb ;for {_ccc ,_gebc =_eae ._gab .ReadBool ();if _gebc !=nil {return _gebc ;};_dfda =_dfda .walk (_ccc );
if _dfda ==nil {continue _cggg ;};if !_dfda ._agd {continue ;};switch _dfda ._gcbf {case _gb :var _abe int ;if _dce {_abe ,_gebc =_eae .decodeRun (_e );}else {_abe ,_gebc =_eae .decodeRun (_ffa );};if _gebc !=nil {return _gebc ;};_bcf +=_abe ;_eae ._cc [_eae ._feb ]=_bcf ;
_eae ._feb ++;if _dce {_abe ,_gebc =_eae .decodeRun (_ffa );}else {_abe ,_gebc =_eae .decodeRun (_e );};if _gebc !=nil {return _gebc ;};_bcf +=_abe ;_eae ._cc [_eae ._feb ]=_bcf ;_eae ._feb ++;case _ce :_gbb :=_eae .getNextChangingElement (_bcf ,_dce )+1;
if _gbb >=_eae ._gga {_bcf =_eae ._ddgd ;}else {_bcf =_eae ._cged [_gbb ];};default:_edg :=_eae .getNextChangingElement (_bcf ,_dce );if _edg >=_eae ._gga ||_edg ==-1{_bcf =_eae ._ddgd +_dfda ._gcbf ;}else {_bcf =_eae ._cged [_edg ]+_dfda ._gcbf ;};_eae ._cc [_eae ._feb ]=_bcf ;
_eae ._feb ++;_dce =!_dce ;};continue _cggg ;};};return nil ;};func _fff (_aafe []byte ,_ggc int ,_aaac int ,_adga bool )([]byte ,int ){var (_bgab code ;_cfga bool ;);for !_cfga {_bgab ,_aaac ,_cfga =_gabe (_aaac ,_adga );_aafe ,_ggc =_adf (_aafe ,_ggc ,_bgab );
};return _aafe ,_ggc ;};type tree struct{_gadb *treeNode };func (_fcaa *Encoder )encodeG4 (_ega [][]byte )[]byte {_fbd :=make ([][]byte ,len (_ega ));copy (_fbd ,_ega );_fbd =_bbe (_fbd );var _bbdd []byte ;var _efe int ;for _eec :=1;_eec < len (_fbd );
_eec ++{if _fcaa .Rows > 0&&!_fcaa .EndOfBlock &&_eec ==(_fcaa .Rows +1){break ;};var _ffc []byte ;var _dab ,_eag ,_ffg int ;_cff :=_efe ;_ebd :=-1;for _ebd < len (_fbd [_eec ]){_dab =_ece (_fbd [_eec ],_ebd );_eag =_bbg (_fbd [_eec ],_fbd [_eec -1],_ebd );
_ffg =_ece (_fbd [_eec -1],_eag );if _ffg < _dab {_ffc ,_cff =_adf (_ffc ,_cff ,_afd );_ebd =_ffg ;}else {if _a .Abs (float64 (_eag -_dab ))> 3{_ffc ,_cff ,_ebd =_fefc (_fbd [_eec ],_ffc ,_cff ,_ebd ,_dab );}else {_ffc ,_cff =_gea (_ffc ,_cff ,_dab ,_eag );
_ebd =_dab ;};};};_bbdd =_fcaa .appendEncodedRow (_bbdd ,_ffc ,_efe );if _fcaa .EncodedByteAlign {_cff =0;};_efe =_cff %8;};if _fcaa .EndOfBlock {_ded ,_ :=_adgd (_efe );_bbdd =_fcaa .appendEncodedRow (_bbdd ,_ded ,_efe );};return _bbdd ;};type code struct{Code uint16 ;
BitsWritten int ;};func (_gfb *Decoder )decodeG32D ()error {_gfb ._gga =_gfb ._feb ;_gfb ._cc ,_gfb ._cged =_gfb ._cged ,_gfb ._cc ;_aedc :=true ;var (_cfe bool ;_cgga int ;_ecf error ;);_gfb ._feb =0;_ea :for _cgga < _gfb ._ddgd {_de :=_g ._gadb ;for {_cfe ,_ecf =_gfb ._gab .ReadBool ();
if _ecf !=nil {return _ecf ;};_de =_de .walk (_cfe );if _de ==nil {continue _ea ;};if !_de ._agd {continue ;};switch _de ._gcbf {case _gb :var _bgb int ;if _aedc {_bgb ,_ecf =_gfb .decodeRun (_e );}else {_bgb ,_ecf =_gfb .decodeRun (_ffa );};if _ecf !=nil {return _ecf ;
};_cgga +=_bgb ;_gfb ._cc [_gfb ._feb ]=_cgga ;_gfb ._feb ++;if _aedc {_bgb ,_ecf =_gfb .decodeRun (_ffa );}else {_bgb ,_ecf =_gfb .decodeRun (_e );};if _ecf !=nil {return _ecf ;};_cgga +=_bgb ;_gfb ._cc [_gfb ._feb ]=_cgga ;_gfb ._feb ++;case _ce :_cb :=_gfb .getNextChangingElement (_cgga ,_aedc )+1;
if _cb >=_gfb ._gga {_cgga =_gfb ._ddgd ;}else {_cgga =_gfb ._cged [_cb ];};default:_fedc :=_gfb .getNextChangingElement (_cgga ,_aedc );if _fedc >=_gfb ._gga ||_fedc ==-1{_cgga =_gfb ._ddgd +_de ._gcbf ;}else {_cgga =_gfb ._cged [_fedc ]+_de ._gcbf ;};
_gfb ._cc [_gfb ._feb ]=_cgga ;_gfb ._feb ++;_aedc =!_aedc ;};continue _ea ;};};return nil ;};type DecodeOptions struct{Columns int ;Rows int ;K int ;EncodedByteAligned bool ;BlackIsOne bool ;EndOfBlock bool ;EndOfLine bool ;DamagedRowsBeforeError int ;
};func (_bge *Decoder )decodeRowType6 ()error {if _bge ._ebf {_bge ._gab .Align ();};if _bge ._ffba {_bge ._gab .Mark ();_cdb ,_aec :=_bge .tryFetchEOL ();if _aec !=nil {return _aec ;};if _cdb {_cdb ,_aec =_bge .tryFetchEOL ();if _aec !=nil {return _aec ;
};if _cdb {return _c .EOF ;};};_bge ._gab .Reset ();};return _bge .decode2D ();};type treeNode struct{_ggbe *treeNode ;_aad *treeNode ;_gcbf int ;_cgeff bool ;_agd bool ;};var _afg =[...][]uint16 {{0x7,0x8,0xb,0xc,0xe,0xf},{0x12,0x13,0x14,0x1b,0x7,0x8},{0x17,0x18,0x2a,0x2b,0x3,0x34,0x35,0x7,0x8},{0x13,0x17,0x18,0x24,0x27,0x28,0x2b,0x3,0x37,0x4,0x8,0xc},{0x12,0x13,0x14,0x15,0x16,0x17,0x1a,0x1b,0x2,0x24,0x25,0x28,0x29,0x2a,0x2b,0x2c,0x2d,0x3,0x32,0x33,0x34,0x35,0x36,0x37,0x4,0x4a,0x4b,0x5,0x52,0x53,0x54,0x55,0x58,0x59,0x5a,0x5b,0x64,0x65,0x67,0x68,0xa,0xb},{0x98,0x99,0x9a,0x9b,0xcc,0xcd,0xd2,0xd3,0xd4,0xd5,0xd6,0xd7,0xd8,0xd9,0xda,0xdb},{},{0x8,0xc,0xd},{0x12,0x13,0x14,0x15,0x16,0x17,0x1c,0x1d,0x1e,0x1f}};
func (_ccf *Decoder )decodeRowType2 ()error {if _ccf ._ebf {_ccf ._gab .Align ();};if _bca :=_ccf .decode1D ();_bca !=nil {return _bca ;};return nil ;};func _fefc (_cbd ,_gfbb []byte ,_bgff ,_ddf ,_gef int )([]byte ,int ,int ){_cbg :=_ece (_cbd ,_gef );
_bcaf :=_ddf >=0&&_cbd [_ddf ]==_dcb ||_ddf ==-1;_gfbb ,_bgff =_adf (_gfbb ,_bgff ,_bbb );var _ccd int ;if _ddf > -1{_ccd =_gef -_ddf ;}else {_ccd =_gef -_ddf -1;};_gfbb ,_bgff =_fff (_gfbb ,_bgff ,_ccd ,_bcaf );_bcaf =!_bcaf ;_fece :=_cbg -_gef ;_gfbb ,_bgff =_fff (_gfbb ,_bgff ,_fece ,_bcaf );
_ddf =_cbg ;return _gfbb ,_bgff ,_ddf ;};var (_dcb byte =1;_cdbf byte =0;);func _bbe (_gefd [][]byte )[][]byte {_gff :=make ([]byte ,len (_gefd [0]));for _bbaf :=range _gff {_gff [_bbaf ]=_dcb ;};_gefd =append (_gefd ,[]byte {});for _deda :=len (_gefd )-1;
_deda > 0;_deda --{_gefd [_deda ]=_gefd [_deda -1];};_gefd [0]=_gff ;return _gefd ;};func (_dcfd *treeNode )walk (_gabc bool )*treeNode {if _gabc {return _dcfd ._aad ;};return _dcfd ._ggbe ;};func _dac (_cca []byte ,_cdee int )([]byte ,int ){return _adf (_cca ,_cdee ,_afd )};
func (_ebfd *Decoder )decodeRun (_cae *tree )(int ,error ){var _edf int ;_fac :=_cae ._gadb ;for {_bac ,_bcfc :=_ebfd ._gab .ReadBool ();if _bcfc !=nil {return 0,_bcfc ;};_fac =_fac .walk (_bac );if _fac ==nil {return 0,_f .New ("\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006f\u0064\u0065\u0020\u0069n\u0020H\u0075f\u0066m\u0061\u006e\u0020\u0052\u004c\u0045\u0020\u0073\u0074\u0072\u0065\u0061\u006d");
};if _fac ._agd {_edf +=_fac ._gcbf ;switch {case _fac ._gcbf >=64:_fac =_cae ._gadb ;case _fac ._gcbf >=0:return _edf ,nil ;default:return _ebfd ._ddgd ,nil ;};};};};func _adf (_ebe []byte ,_dca int ,_efd code )([]byte ,int ){_gfgd :=0;for _gfgd < _efd .BitsWritten {_bcc :=_dca /8;
_gcgc :=_dca %8;if _bcc >=len (_ebe ){_ebe =append (_ebe ,0);};_babcc :=8-_gcgc ;_bbbe :=_efd .BitsWritten -_gfgd ;if _babcc > _bbbe {_babcc =_bbbe ;};if _gfgd < 8{_ebe [_bcc ]=_ebe [_bcc ]|byte (_efd .Code >>uint (8+_gcgc -_gfgd ))&_agb [8-_babcc -_gcgc ];
}else {_ebe [_bcc ]=_ebe [_bcc ]|(byte (_efd .Code <<uint (_gfgd -8))&_agb [8-_babcc ])>>uint (_gcgc );};_dca +=_babcc ;_gfgd +=_babcc ;};return _ebe ,_dca ;};func _gea (_fbdc []byte ,_feg ,_acab ,_gda int )([]byte ,int ){_gdd :=_acb (_acab ,_gda );_fbdc ,_feg =_adf (_fbdc ,_feg ,_gdd );
return _fbdc ,_feg ;};func _bgc (_fdgg ,_fecd []byte ,_agfe int ,_dccc bool )int {_bdg :=_ece (_fecd ,_agfe );if _bdg < len (_fecd )&&(_agfe ==-1&&_fecd [_bdg ]==_dcb ||_agfe >=0&&_agfe < len (_fdgg )&&_fdgg [_agfe ]==_fecd [_bdg ]||_agfe >=len (_fdgg )&&_dccc &&_fecd [_bdg ]==_dcb ||_agfe >=len (_fdgg )&&!_dccc &&_fecd [_bdg ]==_cdbf ){_bdg =_ece (_fecd ,_bdg );
};return _bdg ;};func (_aeb *Decoder )decodeRow ()(_gcce error ){if !_aeb ._ffba &&_aeb ._aea > 0&&_aeb ._aea ==_aeb ._abc {return _c .EOF ;};switch _aeb ._ca {case _eed :_gcce =_aeb .decodeRowType2 ();case _fca :_gcce =_aeb .decodeRowType4 ();case _eee :_gcce =_aeb .decodeRowType6 ();
};if _gcce !=nil {return _gcce ;};_cde :=0;_dbb :=true ;_aeb ._ge =0;for _gd :=0;_gd < _aeb ._feb ;_gd ++{_aed :=_aeb ._ddgd ;if _gd !=_aeb ._feb {_aed =_aeb ._cc [_gd ];};if _aed > _aeb ._ddgd {_aed =_aeb ._ddgd ;};_ac :=_cde /8;for _cde %8!=0&&_aed -_cde > 0{var _fgd byte ;
if !_dbb {_fgd =1<<uint (7-(_cde %8));};_aeb ._cge [_ac ]|=_fgd ;_cde ++;};if _cde %8==0{_ac =_cde /8;var _aca byte ;if !_dbb {_aca =0xff;};for _aed -_cde > 7{_aeb ._cge [_ac ]=_aca ;_cde +=8;_ac ++;};};for _aed -_cde > 0{if _cde %8==0{_aeb ._cge [_ac ]=0;
};var _gcg byte ;if !_dbb {_gcg =1<<uint (7-(_cde %8));};_aeb ._cge [_ac ]|=_gcg ;_cde ++;};_dbb =!_dbb ;};if _cde !=_aeb ._ddgd {return _f .New ("\u0073\u0075\u006d\u0020\u006f\u0066 \u0072\u0075\u006e\u002d\u006c\u0065\u006e\u0067\u0074\u0068\u0073\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074 \u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0063\u0061\u006e\u0020\u006c\u0069\u006ee\u0020w\u0069\u0064\u0074\u0068");
};_aeb ._cgc =(_cde +7)/8;_aeb ._abc ++;return nil ;};func (_edc *Decoder )tryFetchEOL ()(bool ,error ){_facb ,_agbd :=_edc ._gab .ReadBits (12);if _agbd !=nil {return false ,_agbd ;};return _facb ==0x1,nil ;};func (_fad *Decoder )getNextChangingElement (_bd int ,_aac bool )int {_dbg :=0;
if !_aac {_dbg =1;};_cfd :=int (uint32 (_fad ._ge )&0xFFFFFFFE)+_dbg ;if _cfd > 2{_cfd -=2;};if _bd ==0{return _cfd ;};for _cgf :=_cfd ;_cgf < _fad ._gga ;_cgf +=2{if _bd < _fad ._cged [_cgf ]{_fad ._ge =_cgf ;return _cgf ;};};return -1;};func (_fef *Decoder )decodeRowType4 ()error {if !_fef ._dc {return _fef .decoderRowType41D ();
};if _fef ._ebf {_fef ._gab .Align ();};_fef ._gab .Mark ();_ggb ,_bga :=_fef .tryFetchEOL ();if _bga !=nil {return _bga ;};if !_ggb &&_fef ._gbf {_fef ._aae ++;if _fef ._aae > _fef ._cea {return _bg ;};_fef ._gab .Reset ();};if !_ggb {_fef ._gab .Reset ();
};_ecb ,_bga :=_fef ._gab .ReadBool ();if _bga !=nil {return _bga ;};if _ecb {if _ggb &&_fef ._ffba {if _bga =_fef .tryFetchRTC2D ();_bga !=nil {return _bga ;};};_bga =_fef .decode1D ();}else {_bga =_fef .decode2D ();};if _bga !=nil {return _bga ;};return nil ;
};func (_gccb *Decoder )decoderRowType41D ()error {if _gccb ._ebf {_gccb ._gab .Align ();};_gccb ._gab .Mark ();var (_ceb bool ;_gcb error ;);if _gccb ._gbf {_ceb ,_gcb =_gccb .tryFetchEOL ();if _gcb !=nil {return _gcb ;};if !_ceb {return _bg ;};}else {_ceb ,_gcb =_gccb .looseFetchEOL ();
if _gcb !=nil {return _gcb ;};};if !_ceb {_gccb ._gab .Reset ();};if _ceb &&_gccb ._ffba {_gccb ._gab .Mark ();for _ada :=0;_ada < 5;_ada ++{_ceb ,_gcb =_gccb .tryFetchEOL ();if _gcb !=nil {if _f .Is (_gcb ,_c .EOF ){if _ada ==0{break ;};return _ae ;};
};if _ceb {continue ;};if _ada > 0{return _ae ;};break ;};if _ceb {return _c .EOF ;};_gccb ._gab .Reset ();};if _gcb =_gccb .decode1D ();_gcb !=nil {return _gcb ;};return nil ;};var (_eeb map[int ]code ;_cefc map[int ]code ;_eeg map[int ]code ;_db map[int ]code ;
_egf map[int ]code ;_agb map[int ]byte ;_fgg =code {Code :1<<4,BitsWritten :12};_cd =code {Code :3<<3,BitsWritten :13};_ddg =code {Code :2<<3,BitsWritten :13};_afd =code {Code :1<<12,BitsWritten :4};_bbb =code {Code :1<<13,BitsWritten :3};_cec =code {Code :1<<15,BitsWritten :1};
_gf =code {Code :3<<13,BitsWritten :3};_be =code {Code :3<<10,BitsWritten :6};_cg =code {Code :3<<9,BitsWritten :7};_ed =code {Code :2<<13,BitsWritten :3};_bab =code {Code :2<<10,BitsWritten :6};_ga =code {Code :2<<9,BitsWritten :7};);var (_cf *treeNode ;
_ffb *treeNode ;_ffa *tree ;_e *tree ;_fg *tree ;_g *tree ;_ad =-2000;_gc =-1000;_ce =-3000;_gb =-4000;);func (_cac *Encoder )Encode (pixels [][]byte )[]byte {if _cac .BlackIs1 {_dcb =0;_cdbf =1;}else {_dcb =1;_cdbf =0;};if _cac .K ==0{return _cac .encodeG31D (pixels );
};if _cac .K > 0{return _cac .encodeG32D (pixels );};if _cac .K < 0{return _cac .encodeG4 (pixels );};return nil ;};func (_aaa *Decoder )Read (in []byte )(int ,error ){if _aaa ._dfd !=nil {return 0,_aaa ._dfd ;};_geb :=len (in );var (_gccg int ;_cdd int ;
);for _geb !=0{if _aaa ._gg >=_aaa ._cgc {if _ebg :=_aaa .fetch ();_ebg !=nil {_aaa ._dfd =_ebg ;return 0,_ebg ;};};if _aaa ._cgc ==-1{return _gccg ,_c .EOF ;};switch {case _geb <=_aaa ._cgc -_aaa ._gg :_aab :=_aaa ._cge [_aaa ._gg :_aaa ._gg +_geb ];for _ ,_aaed :=range _aab {if !_aaa ._cda {_aaed =^_aaed ;
};in [_cdd ]=_aaed ;_cdd ++;};_gccg +=len (_aab );_aaa ._gg +=len (_aab );return _gccg ,nil ;default:_cgg :=_aaa ._cge [_aaa ._gg :];for _ ,_afb :=range _cgg {if !_aaa ._cda {_afb =^_afb ;};in [_cdd ]=_afb ;_cdd ++;};_gccg +=len (_cgg );_aaa ._gg +=len (_cgg );
_geb -=len (_cgg );};};return _gccg ,nil ;};