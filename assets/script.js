function updateQrCode() {
  let text = document.getElementById("inputText").value;
  let size = document.getElementById("size").value;
  let bgColor = document.getElementById("bgColor").value
  let fgColor = document.getElementById("fgColor").value
  let level = document.getElementById("level").value;
  let disableBorder = document.getElementById("disableBorder").checked;
  generateQrCode(text, size, bgColor, fgColor, level, disableBorder);
}

function generateQrCode(text, size, bgColor, fgColor, level, disableBorder) {
  let b64 = qrCode64(text, size, bgColor, fgColor, level, disableBorder);
  document.getElementById("qrcode").src = `data:image/png;base64,${b64}`;
}

function downloadImage(fileName) {
  const linkSource = document.getElementById("qrcode").src;
  const downloadLink = document.createElement("a");
  
  // Encoding the file name to ensure proper handling on mobile devices
  const encodedFileName = encodeURIComponent(fileName);
  
  downloadLink.href = linkSource;
  downloadLink.download = encodedFileName;
  
  // Append the link to the document body, so it's part of the DOM.
  document.body.appendChild(downloadLink);
  
  // Trigger a click event on the link after a slight delay.
  setTimeout(() => {
      downloadLink.click();
      
      // Clean up by removing the link from the DOM after the click event.
      document.body.removeChild(downloadLink);
  }, 100);
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
