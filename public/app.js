(function (d) {
  const inputFile = d.querySelector('#inputFile');
  const divNotification = d.querySelector('#alert');

  inputFile.addEventListener('change', addFile);

  function addFile(e) {
    const [file] = e.target.files;
    if (!file) {
      return;
    }
    upload(file);
  }

  async function upload(file) {
    const formData = new FormData();
    formData.append('file', file);

    const result = await fetch('/upload', {
      method: 'POST',
      body: formData,
    });

    const message = await result.text();
    onResponse(result.status, message);
  }

  function onResponse(status, message) {
    const className = status !== 400 ? 'success' : 'error';

    divNotification.innerHTML = message;
    divNotification.classList.add(className);
    setTimeout(function () {
      divNotification.classList.remove(className);
    }, 3000);
  }
})(document);
