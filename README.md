# justhost-lookingglass
JustHost Looking Glass

## Usage
```bash
-target string
    Target IP or Hostname
-cookie string
    Cookie
```

If you want to get cookie automatically, you can run this command to install playwright and just leave the cookie empty.
```bash
go run github.com/playwright-community/playwright-go/cmd/playwright install --with-deps
```

## Example
### input cookie manually

```bash
go run . -target 114.114.114.114 -cookie "(your cookie)"
```

### get cookie automatically

```bash
go run . -target 114.114.114.114
```

### output

```bash
justhost-lookingglass % go run . -target 114.114.114.114


<======Helsinki (Tier-3)======>
source(91.149.255.106 / 2605:e440:7::63) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3005ms
rtt min/avg/max/mdev = 121.269/121.330/121.381/0.039 ms


<======IQ Data (Tier-3)======>
source(194.87.68.192 / 2a00:b700:5::5:33c) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3004ms
rtt min/avg/max/mdev = 130.878/130.932/130.957/0.031 ms


<======Amsterdam (Tier-3)======>
source(91.149.241.62 / 2605:e440:2::2e) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3003ms
rtt min/avg/max/mdev = 94.074/96.428/101.850/3.186 ms


<======DallasColo (Tier-3)======>
source(209.209.112.37 / 2605:e440::5:5) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3005ms
rtt min/avg/max/mdev = 22.366/22.416/22.483/0.044 ms


<======Rostelecom======>
source(176.32.32.133 / 2a00:b700:1::4:a) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3004ms
rtt min/avg/max/mdev = 160.816/160.866/160.972/0.062 ms


<======London (Tier-3)======>
source(91.149.202.91 / 2605:e440:8::2:1b) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3005ms
rtt min/avg/max/mdev = 85.133/85.291/85.505/0.144 ms


<======Paris (Tier-3)======>
source(185.143.222.48 / 2605:e440:9::3:20) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3001ms
rtt min/avg/max/mdev = 91.270/91.288/91.309/0.015 ms


<======Madrid (Tier-3)======>
source(91.149.242.38 / 2605:e440:3::2:18) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3004ms
rtt min/avg/max/mdev = 93.938/94.073/94.170/0.088 ms


<======Riga (Tier-3)======>
source(194.58.34.42 / 2605:e440:13::20) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3006ms
rtt min/avg/max/mdev = 123.689/123.781/124.012/0.134 ms


<======Frankfurt (Tier-3)======>
source(91.149.223.204 / 2605:e440:1::1:14a) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3003ms
rtt min/avg/max/mdev = 100.790/100.864/100.944/0.058 ms


<======Los Angeles (Tier-3)======>
source(155.254.192.138 / 2605:e440:5::2:118) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3002ms
rtt min/avg/max/mdev = 56.488/56.506/56.542/0.021 ms


<======TTK======>
source(193.38.50.142 / 2a00:b700:4::1:291) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3005ms
rtt min/avg/max/mdev = 174.407/174.467/174.570/0.061 ms


<======Atlanta (Tier-3)======>
source(134.195.96.22 / 2605:e440:14::3:d) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3005ms
rtt min/avg/max/mdev = 17.644/17.682/17.739/0.038 ms


<======Warsaw (Tier-3)======>
source(91.149.253.58 / 2605:e440:4::2:1f) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3004ms
rtt min/avg/max/mdev = 114.186/114.277/114.487/0.122 ms


<======Kiev (Tier-3)======>
source(194.87.189.8 / ) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3005ms
rtt min/avg/max/mdev = 128.792/128.878/128.964/0.067 ms


<======Fiord======>
source(212.60.5.88 / 2a00:b700:3::42) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3005ms
rtt min/avg/max/mdev = 133.645/133.695/133.743/0.035 ms


<======Toronto (Tier-3)======>
source(104.166.127.78 / 2605:e440:6::71) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3006ms
rtt min/avg/max/mdev = 13.556/13.640/13.782/0.093 ms


<======Rostelecom (ex Adman)======>
source(46.29.162.89 / 2a00:b700:2::6:18) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3005ms
rtt min/avg/max/mdev = 178.207/178.347/178.436/0.088 ms


<======Palermo (Tier-3)======>
source(194.87.47.32 / 2605:e440:12::2:16) ping target(114.114.114.114)
4 packets transmitted, 4 received, 0% packet loss, time 3004ms
rtt min/avg/max/mdev = 134.071/134.148/134.280/0.079 ms


<======Hong Kong (Tier-3)======>
source(91.149.237.213 / 2605:e440:15::2:7b) ping target(114.114.114.114)
5 packets transmitted, 4 received, 20% packet loss, time 4012ms
rtt min/avg/max/mdev = 182.084/185.569/190.964/3.368 ms


<======DataLine (Tier-3)======>
source(46.17.40.85 / 2a00:b700::d:75) ping target(114.114.114.114)
15 packets transmitted, 0 received, 100% packet loss, time 14293ms
```