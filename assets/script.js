function updateQrCode() {
  let text = document.getElementById("inputText").value;
  generateQrCode(text);
}

function generateQrCode(text) {
  let size = document.getElementById("size").value;
  let b64 = qrCode64(size, text);
  document.getElementById("qrcode").src = `data:image/png;base64,${b64}`;
}

function downloadImage(fileName) {
  const linkSource = document.getElementById("qrcode").src;
  const downloadLink = document.createElement("a");
  downloadLink.href = linkSource;
  downloadLink.download = fileName;
  downloadLink.click();
}
function simulateUserInput(inputId, text, interval = 50) {
  let index = 0;

  let inputElement = document.getElementById(inputId);

  const inputInterval = setInterval(() => {
    // Check if the inputElement is valid
    if (
      !inputElement ||
      !inputElement.tagName ||
      inputElement.tagName.toLowerCase() !== "input"
    ) {
      console.error("Invalid input element provided");
      clearInterval(inputInterval);
      return;
    }

    // Set the value of the input element with the next character
    inputElement.value += text[index++];

    // Trigger the 'input' event to simulate user input
    const event = new Event("input", { bubbles: true });
    inputElement.dispatchEvent(event);

    // Stop when all characters have been typed
    if (index >= text.length) {
      clearInterval(inputInterval);
    }
  }, interval);
}

// PWA
if ("serviceWorker" in navigator) {
  window.addEventListener("load", function () {
    navigator.serviceWorker.register("sw.js").then(
      function (registration) {
        console.log(
          "ServiceWorker registration successful with scope: ",
          registration.scope
        );
      },
      function (err) {
        console.log("ServiceWorker registration failed: ", err);
      }
    );
  });
}
