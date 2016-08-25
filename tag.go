package dump

import (
	"fmt"
)

type Tag int

const (
	tagBeginString            Tag = 8
	tagBodyLength             Tag = 9
	tagMsgType                Tag = 35
	tagSenderCompID           Tag = 49
	tagTargetCompID           Tag = 56
	tagOnBehalfOfCompID       Tag = 115
	tagDeliverToCompID        Tag = 128
	tagSecureDataLen          Tag = 90
	tagMsgSeqNum              Tag = 34
	tagSenderSubID            Tag = 50
	tagSenderLocationID       Tag = 142
	tagTargetSubID            Tag = 57
	tagTargetLocationID       Tag = 143
	tagOnBehalfOfSubID        Tag = 116
	tagOnBehalfOfLocationID   Tag = 144
	tagDeliverToSubID         Tag = 129
	tagDeliverToLocationID    Tag = 145
	tagPossDupFlag            Tag = 43
	tagPossResend             Tag = 97
	tagSendingTime            Tag = 52
	tagOrigSendingTime        Tag = 122
	tagXMLDataLen             Tag = 212
	tagXMLData                Tag = 213
	tagMessageEncoding        Tag = 347
	tagLastMsgSeqNumProcessed Tag = 369
	tagOnBehalfOfSendingTime  Tag = 370
	tagApplVerID              Tag = 1128
	tagCstmApplVerID          Tag = 1129
	tagNoHops                 Tag = 627
	tagApplExtID              Tag = 1156
	tagSecureData             Tag = 91
	tagHopCompID              Tag = 628
	tagHopSendingTime         Tag = 629
	tagHopRefID               Tag = 630

	tagHeartBtInt           Tag = 108
	tagBusinessRejectReason Tag = 380
	tagSessionRejectReason  Tag = 373
	tagRefMsgType           Tag = 372
	tagRefTagID             Tag = 371
	tagRefSeqNum            Tag = 45
	tagEncryptMethod        Tag = 98
	tagResetSeqNumFlag      Tag = 141
	tagDefaultApplVerID     Tag = 1137
	tagText                 Tag = 58
	tagTestReqID            Tag = 112
	tagGapFillFlag          Tag = 123
	tagNewSeqNo             Tag = 36
	tagBeginSeqNo           Tag = 7
	tagEndSeqNo             Tag = 16

	tagSignatureLength Tag = 93
	tagSignature       Tag = 89
	tagCheckSum        Tag = 10
)

func (t Tag) String() string {
	switch t {
	case 8:
		return "8(tagBeginString)"
	case 9:
		return "9(tagBodyLength)"
	case 35:
		return "35(tagMsgType)"
	case 49:
		return "49(tagSenderCompID)"
	case 56:
		return "56(tagTargetCompID)"
	case 115:
		return "115(tagOnBehalfOfCompID)"
	case 128:
		return "128(tagDeliverToCompID)"
	case 90:
		return "90(tagSecureDataLen)"
	case 34:
		return "34(tagMsgSeqNum)"
	case 50:
		return "50(tagSenderSubID)"
	case 142:
		return "142(tagSenderLocationID)"
	case 57:
		return "57(tagTargetSubID)"
	case 143:
		return "143(tagTargetLocationID)"
	case 116:
		return "116(tagOnBehalfOfSubID)"
	case 144:
		return "144(tagOnBehalfOfLocationID)"
	case 129:
		return "129(tagDeliverToSubID)"
	case 145:
		return "145(tagDeliverToLocationID)"
	case 43:
		return "43(tagPossDupFlag)"
	case 97:
		return "97(tagPossResend)"
	case 52:
		return "52(tagSendingTime)"
	case 122:
		return "122(tagOrigSendingTime)"
	case 212:
		return "212(tagXMLDataLen)"
	case 213:
		return "213(tagXMLData)"
	case 347:
		return "347(tagMessageEncoding)"
	case 369:
		return "369(tagLastMsgSeqNumProcessed)"
	case 370:
		return "370(tagOnBehalfOfSendingTime)"
	case 1128:
		return "1128(tagApplVerID)"
	case 1129:
		return "1129(tagCstmApplVerID)"
	case 627:
		return "627(tagNoHops)"
	case 1156:
		return "1156(tagApplExtID)"
	case 91:
		return "91(tagSecureData)"
	case 628:
		return "628(tagHopCompID)"
	case 629:
		return "629(tagHopSendingTime)"
	case 630:
		return "630(tagHopRefID)"
	case 108:
		return "108(tagHeartBtInt)"
	case 380:
		return "380(tagBusinessRejectReason)"
	case 373:
		return "373(tagSessionRejectReason)"
	case 372:
		return "372(tagRefMsgType)"
	case 371:
		return "371(tagRefTagID)"
	case 45:
		return "45(tagRefSeqNum)"
	case 98:
		return "98(tagEncryptMethod)"
	case 141:
		return "141(tagResetSeqNumFlag)"
	case 1137:
		return "1137(tagDefaultApplVerID)"
	case 58:
		return "58(tagText)"
	case 112:
		return "112(tagTestReqID)"
	case 123:
		return "123(tagGapFillFlag)"
	case 36:
		return "36(tagNewSeqNo)"
	case 7:
		return "7(tagBeginSeqNo)"
	case 16:
		return "16(tagEndSeqNo)"
	case 93:
		return "93(tagSignatureLength)"
	case 89:
		return "89(tagSignature)"
	case 10:
		return "10(tagCheckSum)"
	default:
		return fmt.Sprintf("%d", t)
	}
}
