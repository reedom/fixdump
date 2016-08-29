fixdump
=======

Dump FIX protocol log in _easy-read-formats_.

Install
-------

    go get github.com/reedom/fixdump

Usage
-----

    Usage:
      fixdump [OPTIONS] [Paths...]

    Application Options:
      -v, --version
      -H, --human    Additionally print tag/value names
      -i, --indent   Indent tag entries
      --no-heartbeats  Omit heartbeat log from output

    Help Options:
      -h, --help     Show this help message

    Arguments:
      Paths:         Log file path(s) to read

Example Usage
-------------

```
# read from file with no options
fixdump sample.log

[2016/08/27 08:41:27.251195 (2) TW -> ISLD]
8=FIX.4.2
9=104
35=D
34=2
49=TW
52=20140515-19:49:56.659
56=ISLD
11=100
21=1
40=1
54=1
55=TSLA
60=00010101-00:00:00.000
10=039
[2016/08/27 08:41:39.666251 (3) ISLD -> TW]
8=FIX.4.2
…
```

```
# read from STDIN with indent and human options
tail -f sample.log | fixdump -H -i

[2016/08/27 08:41:27.251195 (2) TW -> ISLD]
  8(BeginString)=FIX.4.2
  9(BodyLength)=104
  35(MsgType)=D(Order Single)
  34(MsgSeqNum)=2
  49(SenderCompID)=TW
  52(SendingTime)=20140515-19:49:56.659
  56(TargetCompID)=ISLD
  11(ClOrdID)=100
  21(HandlInst)=1(Automated execution order, private, no Broker intervention)
  40(OrdType)=1(Market)
  54(Side)=1(Buy)
  55(Symbol)=TSLA
  60(TransactTime)=00010101-00:00:00.000
  10(CheckSum)=039
[2016/08/27 08:41:39.666251 (3) ISLD -> TW]
  8(BeginString)=FIX.4.2
…
```

Dump Format
-----------

__Header Line__

```
[2016/08/27 08:41:27.251195 (2) TW -> ISLD]
          │                  │  │       └── TargetCompID
          │                  │  └── SenderCompID
          │                  └── MsgSeqNum
          └─ timestamp
```

__Field Line__

```
  35(MsgType)=D(Order Single)
   │     │    │    └── meaning of the value
   │     │    └── value
   │     └── field name
   └─ tag


```

licence
-------

MIT

