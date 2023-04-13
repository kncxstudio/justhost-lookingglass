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

