import { getDevice } from "./helper.js";

function main(iosLink, androidLink) {
  const device = getDevice();
  switch (device) {
    case "android":
      try {
        window.location = androidLink;
      } catch (e) {
        console.log(`found error\n${e}`);
      }
      setTimeout(() => (window.location = "/store"), 250);
      break;

    case "iphone":
      try {
        window.location = iosLink;
      } catch (e) {
        console.log(`found error\n${e}`);
      }
      setTimeout(() => (window.location = "/store"), 250);
      break;
  }
}

window.main = main;
