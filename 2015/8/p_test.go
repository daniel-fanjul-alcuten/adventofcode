package p

import "testing"

func TestP1_S1(t *testing.T) {
	n1, n2 := P1([]string{`""`})
	if n1 != 2 {
		t.Error(n1)
	}
	if n2 != 0 {
		t.Error(n2)
	}
}

func TestP1_S2(t *testing.T) {
	n1, n2 := P1([]string{`"abc"`})
	if n1 != 5 {
		t.Error(n1)
	}
	if n2 != 3 {
		t.Error(n2)
	}
}

func TestP1_S3(t *testing.T) {
	n1, n2 := P1([]string{`"aaa\"aaa"`})
	if n1 != 10 {
		t.Error(n1)
	}
	if n2 != 7 {
		t.Error(n2)
	}
}

func TestP1_S4(t *testing.T) {
	n1, n2 := P1([]string{`"\x27"`})
	if n1 != 6 {
		t.Error(n1)
	}
	if n2 != 1 {
		t.Error(n2)
	}
}

func TestP1_X(t *testing.T) {
	n1, n2 := P1(X)
	if n1 != 6202 {
		t.Error(n1)
	}
	if n2 != 4860 {
		t.Error(n2)
	}
	if n := n1 - n2; n != 1342 {
		t.Error(n)
	}
}

func TestP2_S1(t *testing.T) {
	n1, n2 := P2([]string{`""`})
	if n1 != 6 {
		t.Error(n1)
	}
	if n2 != 2 {
		t.Error(n2)
	}
}

func TestP2_S2(t *testing.T) {
	n1, n2 := P2([]string{`"abc"`})
	if n1 != 9 {
		t.Error(n1)
	}
	if n2 != 5 {
		t.Error(n2)
	}
}

func TestP2_S3(t *testing.T) {
	n1, n2 := P2([]string{`"aaa\"aaa"`})
	if n1 != 16 {
		t.Error(n1)
	}
	if n2 != 10 {
		t.Error(n2)
	}
}

func TestP2_S4(t *testing.T) {
	n1, n2 := P2([]string{`"\x27"`})
	if n1 != 11 {
		t.Error(n1)
	}
	if n2 != 6 {
		t.Error(n2)
	}
}

func TestP2_X(t *testing.T) {
	n1, n2 := P2(X)
	if n1 != 8276 {
		t.Error(n1)
	}
	if n2 != 6202 {
		t.Error(n2)
	}
	if n := n1 - n2; n != 2074 {
		t.Error(n)
	}
}

var X = []string{
	`"azlgxdbljwygyttzkfwuxv"`,
	`"v\xfb\"lgs\"kvjfywmut\x9cr"`,
	`"merxdhj"`,
	`"dwz"`,
	`"d\\gkbqo\\fwukyxab\"u"`,
	`"k\xd4cfixejvkicryipucwurq\x7eq"`,
	`"nvtidemacj\"hppfopvpr"`,
	`"kbngyfvvsdismznhar\\p\"\"gpryt\"jaeh"`,
	`"khre\"o\x0elqfrbktzn"`,
	`"nugkdmqwdq\x50amallrskmrxoyo"`,
	`"jcrkptrsasjp\\\"cwigzynjgspxxv\\vyb"`,
	`"ramf\"skhcmenhbpujbqwkltmplxygfcy"`,
	`"aqjqgbfqaxga\\fkdcahlfi\"pvods"`,
	`"pcrtfb"`,
	`"\x83qg\"nwgugfmfpzlrvty\"ryoxm"`,
	`"fvhvvokdnl\\eap"`,
	`"kugdkrat"`,
	`"seuxwc"`,
	`"vhioftcosshaqtnz"`,
	`"gzkxqrdq\\uko\"mrtst"`,
	`"znjcomvy\x16hhsenmroswr"`,
	`"clowmtra"`,
	`"\xc4"`,
	`"jpavsevmziklydtqqm"`,
	`"egxjqytcttr\\ecfedmmovkyn\"m"`,
	`"mjulrvqgmsvmwf"`,
	`"o\\prxtlfbatxerhev\xf9hcl\x44rzmvklviv"`,
	`"lregjexqaqgwloydxdsc\\o\"dnjfmjcu"`,
	`"lnxluajtk\x8desue\\k\x7abhwokfhh"`,
	`"wrssfvzzn\"llrysjgiu\"npjtdli"`,
	`"\x67lwkks"`,
	`"bifw\"ybvmwiyi\"vhol\"vol\xd4"`,
	`"aywdqhvtvcpvbewtwuyxrix"`,
	`"gc\xd3\"caukdgfdywj"`,
	`"uczy\\fk"`,
	`"bnlxkjvl\x7docehufkj\\\"qoyhag"`,
	`"bidsptalmoicyorbv\\"`,
	`"jorscv\"mufcvvfmcv\"ga"`,
	`"sofpwfal\\a"`,
	`"kcuqtbboaly\"uj\"k"`,
	`"n\\c"`,
	`"x\"\xcaj\\xwwvpdldz"`,
	`"eyukphh"`,
	`"wcyjq"`,
	`"vjx\"\"hjroj\"l\x4cjwbr"`,
	`"xcodsxzfqw\\rowqtuwvjnxupjnrh"`,
	`"yc"`,
	`"fpvzldgbdtca\"hqwa"`,
	`"ymjq\x8ahohvafubra\"hgqoknkuyph"`,
	`"kx\\mkaaklvcup"`,
	`"belddrzegcsxsyfhzyz"`,
	`"fuyswi"`,
	`"\\hubzebo\"ha\\qyr\"dv\\"`,
	`"mxvlz\"fwuvx\"cyk\""`,
	`"ftbh\"ro\\tmcpnpvh\"xx"`,
	`"ygi"`,
	`"rw\"\"wwn\\fgbjumq\"vgvoh\xd0\"mm"`,
	`"\"pat\"\x63kpfc\"\x2ckhfvxk\"uwqzlx"`,
	`"o"`,
	`"d\"hqtsfp\xceaswe\"\xc0lw"`,
	`"zajpvfawqntvoveal\"\"trcdarjua"`,
	`"xzapq"`,
	`"rkmhm"`,
	`"byuq"`,
	`"rwwmt\xe8jg\xc2\"omt"`,
	`"nfljgdmgefvlh\"x"`,
	`"rpjxcexisualz"`,
	`"doxcycmgaiptvd"`,
	`"rq\\\"mohnjdf\\xv\\hrnosdtmvxot"`,
	`"oqvbcenib\"uhy\\npjxg"`,
	`"pkvgnm\\ruayuvpbpd"`,
	`"kknmzpxqfbcdgng"`,
	`"piduhbmaympxdexz"`,
	`"vapczawekhoa\\or"`,
	`"tlwn\"avc\"bycg\"\"xuxea"`,
	`"\xcdvryveteqzxrgopmdmihkcgsuozips"`,
	`"kpzziqt"`,
	`"sdy\\s\"cjq"`,
	`"yujs"`,
	`"qte\"q"`,
	`"qyvpnkhjcqjv\"cclvv\"pclgtg\xeak\"tno"`,
	`"xwx"`,
	`"vibuvv"`,
	`"qq\""`,
	`"wwjduomtbkbdtorhpyalxswisq\"r"`,
	`"afuw\\mfjzctcivwesutxbk\"lk"`,
	`"e\xcef\\hkiu"`,
	`"ftdrgzvygcw\"jwsrcmgxj"`,
	`"zrddqfkx\x21dr\"ju\"elybk\"powj\"\"kpryz"`,
	`"dttdkfvbodkma\""`,
	`"lzygktugpqw"`,
	`"qu\x83tes\\u\"tnid\"ryuz"`,
	`"\\o\"pe\\vqwlsizjklwrjofg\xe2oau\\rd"`,
	`"mikevjzhnwgx\"fozrj\"h\""`,
	`"ligxmxznzvtachvvbahnff"`,
	`"d\\kq"`,
	`"tnbkxpzmcakqhaa"`,
	`"g\\yeakebeyv"`,
	`"cqkcnd\"sxjxfnawy\x31zax\x6ceha"`,
	`"m\x0dtqotffzdnetujtsgjqgwddc"`,
	`"masnugb\"etgmxul\x3bqd\\tmtddnvcy"`,
	`"floediikodfgre\x23wyoxlswxflwecdjpt"`,
	`"zu"`,
	`"r"`,
	`"\"ashzdbd\"pdvba\xeeumkr\\amnj"`,
	`"ckslmuwbtfouwpfwtuiqmeozgspwnhx"`,
	`"t\\qjsjek\xf9gjcxsyco\"r"`,
	`"hoed\x1b\\tcmaqch\"epdy"`,
	`"mgjiojwzc\\ypqcn\xb1njmp\"aeeblxt"`,
	`"\xdf\"h\x5enfracj"`,
	`"\x6fpbpocrb"`,
	`"jbmhrswyyq\\"`,
	`"wtyqtenfwatji\"ls\\"`,
	`"voy"`,
	`"awj"`,
	`"rtbj\"j"`,
	`"hynl"`,
	`"orqqeuaat\\xu\\havsgr\xc5qdk"`,
	`"g\"npyzjfq\"rjefwsk"`,
	`"rk\\kkcirjbixr\\zelndx\"bsnqvqj\""`,
	`"tecoz"`,
	`"dn\"uswngbdk\""`,
	`"qb\\"`,
	`"wpyis\\ebq"`,
	`"ppwue\\airoxzjjdqbvyurhaabetv"`,
	`"fxlvt"`,
	`"ql\"oqsmsvpxcg\"k"`,
	`"vqlhuec\\adw"`,
	`"qzmi\xffberakqqkk"`,
	`"tisjqff\"wf"`,
	`"yhnpudoaybwucvppj"`,
	`"xhfuf\\ehsrhsnfxcwtibd\"ubfpz"`,
	`"ihgjquzhf\""`,
	`"ff\x66dsupesrnusrtqnywoqcn\\"`,
	`"z\x77zpubbjmd"`,
	`"\"vhzlbwq\"xeimjt\\xe\x85umho\"m\"\"bmy"`,
	`"mmuvkioocmzjjysi\"mkfbec\""`,
	`"rpgghowbduw\x2fayslubajinoik\xd0hcfy"`,
	`"xrkyjqul\xdexlojgdphczp\"jfk"`,
	`"mg\x07cnr\x8b\x67xdgszmgiktpjhawho"`,
	`"kdgufhaoab"`,
	`"rlhela\"nldr"`,
	`"wzye\x87u"`,
	`"yif\x75bjhnitgoarmfgqwpmopu"`,
	`"pvlbyez\"wyy\x3dpgr"`,
	`"ezdm\"ovkruthkvdwtqwr\"ibdoawzgu"`,
	`"qubp"`,
	`"b\\kcpegcn\\zgdemgorjnk"`,
	`"gjsva\\kzaor\"\"gtpd"`,
	`"\"kt"`,
	`"rlymwlcodix"`,
	`"qqtmswowxca\"jvv"`,
	`"jni\xebwhozb"`,
	`"zhino\"kzjtmgxpi\"zzexijg"`,
	`"tyrbat\\mejgzplufxixkyg"`,
	`"lhmopxiao\x09\"p\xebl"`,
	`"xefioorxvate"`,
	`"nmcgd\x46xfujt\"w"`,
	`"\xe3wnwpat\"gtimrb"`,
	`"wpq\"xkjuw\xebbohgcagppb"`,
	`"fmvpwaca"`,
	`"mlsw"`,
	`"fdan\\\x9e"`,
	`"\"f\"fmdlzc"`,
	`"nyuj\\jnnfzdnrqmhvjrahlvzl"`,
	`"zn\"f\xcfsshcdaukkimfwk"`,
	`"uayugezzo\\\"e\"blnrgjaupqhik"`,
	`"efd\"apkndelkuvfvwyyatyttkehc"`,
	`"ufxq\\\"m\"bwkh\x93kapbqrvxxzbzp\\"`,
	`"fgypsbgjak\x79qblbeidavqtddfacq\\i\"h"`,
	`"kcfgpiysdxlgejjvgndb\\dovfpqodw"`,
	`"\"onpqnssmighipuqgwx\"nrokzgvg"`,
	`"vhjrrhfrba\"jebdanzsrdusut\\wbs"`,
	`"o\xdakymbaxakys"`,
	`"uwxhhzz\\mtmhghjn\\\\tnhzbejj"`,
	`"yd\\"`,
	`"bpgztp\\lzwpdqju\"it\x35qjhihjv"`,
	`"\\my\\b\"klnnto\\\xb3mbtsh"`,
	`"ezyvknv\"l\x2bdhhfjcvwzhjgmhwbqd\"\\"`,
	`"ftkz\"amoncbsohtaumhl\"wsodemopodq"`,
	`"ifv"`,
	`"dmzfxvzq"`,
	`"sped\"bvmf\"mmevl\"zydannpfny"`,
	`"fjxcjwlv\"pnqyrzatsjwsqfidb"`,
	`"muc\xfdqouwwnmuixru\\zlhjintplvtee"`,
	`"mraqgvmj"`,
	`"njopq\"ftcsryo"`,
	`"enoh\"n"`,
	`"t\"ntjhjc\"nzqh\xf7dcohhlsja\x7dtr"`,
	`"flbqcmcoun"`,
	`"dxkiysrn\\dyuqoaig"`,
	`"nehkzi\"h\"syktzfufotng\xdafqo"`,
	`"dzkjg\\hqjk\\\"zfegssjhn"`,
	`"sadlsjv"`,
	`"vmfnrdb\""`,
	`"ac\\bdp\"n"`,
	`"qt\x89h"`,
	`"lsndeugwvijwde\\vjapbm\\k\\nljuva"`,
	`"twpmltdzyynqt\\z\\tnund\x64hm"`,
	`"hpcyata\"ocylbkzdnhujh"`,
	`"hskzq\"knntuhscex\"q\\y\\vqj\x3an"`,
	`"eekwyufvji\\mqgeroekxeyrmymq"`,
	`"hl\"durthetvri\xebw\\jxu\"rcmiuy"`,
	`"\"fxdnmvnftxwesmvvq\"sjnf\xaabpg\"iary"`,
	`"\"\"nksqso"`,
	`"ruq\xbezugge\"d\"hwvoxmy\"iawikddxn\"x"`,
	`"rxxnlfay"`,
	`"stcu\"mv\xabcqts\\fasff"`,
	`"yrnvwfkfuzuoysfdzl\x02bk"`,
	`"qbdsmlwdbfknivtwijbwtatqfe"`,
	`"\"erqh\\csjph"`,
	`"ikfv"`,
	`"\xd2cuhowmtsxepzsivsvnvsb"`,
	`"vj"`,
	`"d"`,
	`"\\g"`,
	`"porvg\x62qghorthnc\"\\"`,
	`"tiks\\kr\"\x0fuejvuxzswnwdjscrk"`,
	`"xmgfel\"atma\\zaxmlgfjx\"ajmqf"`,
	`"oz\\rnxwljc\\\"umhymtwh"`,
	`"wlsxxhm\x7fqx\\gjoyrvccfiner\\qloluqv"`,
	`"k\\ieq"`,
	`"xidjj\"ksnlgnwxlddf\\s\\kuuleb"`,
	`"wjpnzgprzv\\maub\x0cj"`,
	`"r"`,
	`"y"`,
	`"\"yecqiei\"ire\\jdhlnnlde\xc5u"`,
	`"drvdiycqib"`,
	`"egnrbefezcrhgldrtb"`,
	`"plqodxv\\zm\"uodwjdocri\x55ucaezutm"`,
	`"f\"wexcw\x02ekewx\"alyzn"`,
	`"pqajwuk\\\\oatkfqdyspnrupo"`,
	`"rkczj\"fzntabpnygrhamk\\km\x68xfkmr"`,
	`"wejam\xbac\x37kns"`,
	`"qqmlwjk\"gh"`,
	`"fdcjsxlgx"`,
	`"\\cxvxy\"kb\"\"unubvrsq\\y\\awfhbmarj\\"`,
	`"geunceaqr"`,
	`"tpkg\"svvngk\\sizlsyaqwf"`,
	`"\"pa\\x\x18od\\emgje\\"`,
	`"ffiizogjjptubzqfuh\"cctieqcdh"`,
	`"yikhiyyrpgglpos"`,
	`"h\\"`,
	`"jotqojodcv"`,
	`"ervsz\x87ade\"fevq\\tcqowt"`,
	`"\\y\"fgrxtppkcseeg\\onxjarx\\hyhfn\x5fi"`,
	`"kxndlabn\\wwumctuzdcfiitrbnn"`,
	`"eoosynwhwm"`,
	`"\"c\x04"`,
	`"ny\xf6vuwlec"`,
	`"ubgxxcvnltzaucrzg\\xcez"`,
	`"pnocjvo\\yt"`,
	`"fcabrtqog\"a\"zj"`,
	`"o\\bha\\mzxmrfltnflv\xea"`,
	`"tbfvzwhexsdxjmxejwqqngzixcx"`,
	`"wdptrakok\"rgymturdmwfiwu"`,
	`"reffmj"`,
	`"lqm"`,
	`"\\oc"`,
	`"p\""`,
	`"ygkdnhcuehlx"`,
	`"vsqmv\"bqay\"olimtkewedzm"`,
	`"isos\x6azbnkojhxoopzetbj\xe1yd"`,
	`"yo\\pgayjcyhshztnbdv"`,
	`"fg\"h"`,
	`"vcmcojolfcf\\\\oxveua"`,
	`"w\"vyszhbrr\"jpeddpnrjlca\x69bdbopd\\z"`,
	`"jikeqv"`,
	`"\"dkjdfrtj"`,
	`"is"`,
	`"hgzx"`,
	`"z\""`,
	`"woubquq\\ag\""`,
	`"xvclriqa\xe6ltt"`,
	`"tfxinifmd"`,
	`"mvywzf\"jz"`,
	`"vlle"`,
	`"c\"rf\"wynhye\x25vccvb\""`,
	`"zvuxm"`,
	`"\xf2\"jdstiwqer\"h"`,
	`"kyogyogcknbzv\x9f\\\\e"`,
	`"kspodj\"edpeqgypc"`,
	`"oh\\x\\h"`,
	`"julb"`,
	`"bmcfkidxyilgoy\\xmu\"ig\\qg"`,
	`"veqww\"ea"`,
	`"fkdbemtgtkpqisrwlxutllxc\"mbelhs"`,
	`"e"`,
	`"ecn\x50ooprbstnq"`,
	`"\"\xe8\"ec\xeah\"qo\\g\"iuqxy\"e\"y\xe7xk\xc6d"`,
	`"lwj\"aftrcqj"`,
	`"jduij\x97zk\"rftjrixzgscxxllpqx\"bwwb"`,
	`"fqcditz"`,
	`"f\x19azclj\"rsvaokgvty\"aeq"`,
	`"erse\x9etmzhlmhy\x67yftoti"`,
	`"lsdw\xb3dmiy\\od"`,
	`"x\x6fxbljsjdgd\xaau"`,
	`"hjg\\w\"\x78uoqbsdikbjxpip\"w\"jnhzec"`,
	`"gk"`,
	`"\\zrs\\syur"`,
}
