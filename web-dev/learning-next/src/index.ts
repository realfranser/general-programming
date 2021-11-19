import * as _ from "lodash";

function component() {
  const element = document.createElement("div");

  element.innerHTML = _.join(["Hello", "next, react and webpack"], " ");

  return element;
}

if (process.env.NODE_ENV !== "production") {
  console.log("Looks like we are in development mode!");
}

document.body.appendChild(component());
