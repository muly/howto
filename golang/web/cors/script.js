
function getData() {

  const url = 'http://localhost:8080'+"?name="+document.getElementById('query').value;
  var newHttp = new XMLHttpRequest();
 
  newHttp.open("GET",url);
  newHttp.send();
  
  newHttp.onreadystatechange = (e) => {
      document.getElementById('output').textContent = newHttp.responseText
  }
}

