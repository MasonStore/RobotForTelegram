#!/bin/zsh

APP_CERTIFICATE="3rd Party Mac Developer Application: Beijing Xingxing Technology Co., Ltd. (6369482DA4)"
PKG_CERTIFICATE="3rd Party Mac Developer Installer: Beijing Xingxing Technology Co., Ltd. (6369482DA4)"
APP_NAME="Robot For Telegram"

# shellcheck disable=SC2164
cd ../app/

rm -rf "./${APP_NAME}.pkg"
rm -rf "./build/bin/${APP_NAME}.app"

wails build -platform darwin/universal -clean -ldflags "-s -w"

cp ./embedded.provisionprofile "./build/bin/$APP_NAME.app/Contents"

chmod -R a+xr "./build/bin/${APP_NAME}.app"

codesign --timestamp --options=runtime -s "$APP_CERTIFICATE" -v --entitlements ./build/darwin/entitlements.plist "./build/bin/${APP_NAME}.app"

productbuild --sign "$PKG_CERTIFICATE" --component "./build/bin/${APP_NAME}.app" /Applications "./${APP_NAME}.pkg"

