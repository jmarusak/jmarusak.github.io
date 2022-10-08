function counter() {
  fetch("http://localhost:8080/counter", {
      method: "GET",
      mode: "cors",
      headers: {}
    })
    .then(response => {
      if (!response.ok) {
        throw new Error(response.error)
      }
      return response.json(); 
    })
    .then(counter => {
      let visitInfo = "visits: " + counter.Visits + " (" + counter.LastVisit + ")"; 
      console.log(visitInfo)
      document.getElementById('counter').innerHTML = visitInfo 
    })
    .catch(function(error) {
      console.log(error)
    });
}
