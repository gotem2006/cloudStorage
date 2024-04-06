function uploadFolder(){
    var folderInput = document.getElementById("folederInput");
    var files = folderInput.files;
    const folderName = files[0].webkitRelativePath.split('/')[0];

    const formData = new FormData();
    formData.append("name", folderName);

    for (var i = 0; i < files.length; i++) {
        formData.append("file", files[i]);
    }


    fetch("http://localhost:8080/ewqwewq/directory", {
        method: "POST",
        body: formData
    })
    .then(response => response.text())
    .then(data => {
        console.log(data);
        alert('Folder uploaded successfully!');
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Error uploading folder. Please try again.');
    });
}