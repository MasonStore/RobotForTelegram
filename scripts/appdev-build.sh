#!/bin/zsh

APP_CERTIFICATE="Mac Developer: ShouYue Ma (VGY8XBV8BF)"
APP_NAME="Robot For Telegram"

# shellcheck disable=SC2164
cd ../app/

rm -rf "./${APP_NAME}.pkg"
rm -rf "./build/bin/${APP_NAME}.app"

wails build -platform darwin/universal -clean -ldflags "-s -w"

cp ./build/darwin/RobotForTelegramDevelopment.provisionprofile "./build/bin/$APP_NAME.app/Contents/embedded.provisionprofile"

chmod -R a+xr "./build/bin/${APP_NAME}.app"

codesign --timestamp --options=runtime -s "$APP_CERTIFICATE" -v --entitlements ./build/darwin/entitlements.plist "./build/bin/${APP_NAME}.app"
