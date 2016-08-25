package logcat

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

const defaults = `
8    BeginString
9    BodyLength
35   MsgType
49   SenderCompID
56   TargetCompID
115  OnBehalfOfCompID
128  DeliverToCompID
90   SecureDataLen
34   MsgSeqNum
50   SenderSubID
142  SenderLocationID
57   TargetSubID
143  TargetLocationID
116  OnBehalfOfSubID
144  OnBehalfOfLocationID
129  DeliverToSubID
145  DeliverToLocationID
43   PossDupFlag
97   PossResend
52   SendingTime
122  OrigSendingTime
212  XMLDataLen
213  XMLData
347  MessageEncoding
369  LastMsgSeqNumProcessed
370  OnBehalfOfSendingTime
1128 ApplVerID
1129 CstmApplVerID
627  NoHops
1156 ApplExtID
91   SecureData
628  HopCompID
629  HopSendingTime
630  HopRefID
108  HeartBtInt
380  BusinessRejectReason
373  SessionRejectReason
372  RefMsgType
371  RefTagID
45   RefSeqNum
98   EncryptMethod
141  ResetSeqNumFlag
1137 DefaultApplVerID
58   Text
112  TestReqID
123  GapFillFlag
36   NewSeqNo
7    BeginSeqNo
16   EndSeqNo
93   SignatureLength
89   Signature
10   CheckSum
`

func (app *App) populateTags() error {
	var reader *bufio.Reader
	if app.opts.TagPath == "" {
		reader = bufio.NewReader(strings.NewReader(defaults))
	} else {
		file, err := os.Open(app.opts.TagPath)
		if err != nil {
			return err
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	}

	re := regexp.MustCompile(`^\s*(\d+)\s+(.*)`)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "TAG file read error: %s\n", err.Error())
			}
			return nil
		}

		parts := re.FindStringSubmatch(line)
		if len(parts) == 3 {
			app.tags[parts[1]] = strings.Trim(parts[2], " \t")
		}
	}
}
