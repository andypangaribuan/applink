export function getDevice() {
  const userAgent = navigator.userAgent;

  if (/android/i.test(userAgent)) {
    return "android";
  } else if (/iphone/i.test(userAgent)) {
    return "iphone";
  } else if (/ipad/i.test(userAgent)) {
    return "ipad";
  }

  return "unknown";
}
