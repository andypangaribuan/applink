import { getDevice } from "./helper.js";

function main(appleStore, googleStore) {
  const device = getDevice();
  switch (device) {
    case "android":
      try {
        window.location = googleStore;
      } catch (e) {
        console.log(`found error\n${e}`);
      }
      break;

    case "iphone":
      try {
        window.location = appleStore;
      } catch (e) {
        console.log(`found error\n${e}`);
      }
      break;
  }
}

window.main = main;
