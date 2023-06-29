let fs = require("fs");

const fixPathInFile = (filePath) => {
  let content = fs.readFileSync(filePath, 'utf8');
  if (content) {
    let replaced = false;
    content = content.replaceAll('www', () => {
      replaced = true;
      return 'public';
    });
    if (replaced) {
      fs.writeFileSync(filePath, content);
    }
  } else {
    console.error('file NativeAudio not found!');
  }
}

if (process.env.CAPACITOR_PLATFORM_NAME === 'android') {
  let filePath = './android/capacitor-cordova-android-plugins/src/main/java/com/rjfun/cordova/plugin/nativeaudio/NativeAudio.java';
  fixPathInFile(filePath);
} else if (process.env.CAPACITOR_PLATFORM_NAME === 'ios') {
  let filePath = './ios/capacitor-cordova-ios-plugins/sources/CordovaPluginNativeaudio/NativeAudio.m';
  fixPathInFile(filePath);
}
