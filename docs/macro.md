# Macro packet structure

600 — 0x258

0x02 0x58

```sh
255 # Number of repetitions

3   # Delay (start of pressing)
12  # Key code 'I'

131 # Release next
12  # Key code 'I'

3   # Delay (start of pressing)
18  # Key code 'O'

3   # Delay (start of pressing)
15  # Key code 'L'

131 # Release next
15  # Key code 'L'

131 # Release next
18  # Key code 'O'
```

If delay more than 64, then format:

```sh
LOW
KEYCODE
0
3
HIGH
MID
```

Formula is (25600 * HIGH) + (MID * 100) + LOW

If LOW is more then 128, it means key up event is sent. So subtract 128 from LOW.

If the LOW rest is 3, then it should be taken as 0 in calculations. 


