// This line waits for the entire page to load before executing the following code
window.onload = function () {

  // This line fetches data from a URL path '/students' (likely an endpoint serving student data)
  fetch('/students')
    .then(response => response.text())  // This part processes the response from the fetch
      //  - The first .then handles the successful fetch. 'response' is the response object.
      //  - We use the response.text() method to convert the response to a text string.

    .then(data => showStudents(data));  // This part handles the processed data
      //  - The second .then receives the text data from the previous step and stores it in 'data'.
      //  - We call a function named showStudent() and pass the 'data' as an argument.
};



function addStudent() {
  var data = getFormData()
  fetch('/student', {
      method :"POST",
      // converts the java scriot format to JSON object formate
      // java script method allows attachment of function in java scritp object.
      // JASON is language indepentdent
      body: JSON.stringify(data),     // converts it into a JSON (JavaScript Object Notation) formatted string
      headers: {"Content-type": "application/json; charset=UTF-8"}
  }).then(response => {  // Assuming this is the success callback of a previous asynchronous operation

      // Extract the student ID from the previously fetched data
      // var sid = data.stdid;
      
      // createing a form validation
      // if (isNaN(sid)) {
      //     alert("Enter a valide student ID")
      //     return
      // } else if (data.email == ""){
      //     alert("Email cannot be empty")
      //     return
      // } else if (data.fname == ""){
      //     alert("First name cannot be empty")
      //     return
      // }
      // Check if the first fetch operation was successful (response1.ok is true)
      if (response.ok) {
        // Chain another fetch request to '/student/' + sid (constructed student URL)
        fetch('/student/' + sid)
          .then(response => response.text())  // Parse the response as text
          .then(data => showStudent(data));    // Call the showStudent function with the fetched student data
      }else {
          throw new Error(response.status)
      }
    }).catch(e => {   // The catch block defines what happens if an error (exception) occurs inside the try block.
      if (e.message == 303) {
          alert("User not logged in.")
          window.open("index.html", "_self")
      } else if (e.message == 500) {
          alert("Server error!")
      }
      // e => { ... }: This part defines an arrow function that will be executed if an error is thrown in
      // alert(e)        // display an alert message. 
      // The argument to alert() is set to e, which is the error object.
    })
    resetform();
}

function showStudent(data) {
  const student = JSON.parse(data)    // The code JSON.parse(data) is used in JavaScript to convert a JSON string (data) into a JavaScript object.
  newRow(student)
}

// sets the form fields to empty 
function resetform(){
  document.getElementById("sid").value = "";
  document.getElementById("fname").value = "";
  document.getElementById("lname").value = "";
  document.getElementById("email").value = "";
}

function showStudents(data) {
  const students = JSON.parse(data)    // The code JSON.parse(data) is used in JavaScript to convert a JSON string (data) into a JavaScript object.
  students.forEach(stud => {
      newRow(stud)
  })

}

var selectedRow = null;
function updateStudent(r){
  selectedRow = r.parentElement.parentElement;
  // fill in the form fields with selected row data
  document.getElementById("sid").value = selectedRow.cells[0].innerHTML;
  document.getElementById("fname").value = selectedRow.cells[1].innerHTML;
  document.getElementById("lname").value = selectedRow.cells[2].innerHTML;
  document.getElementById("email").value = selectedRow.cells[3].innerHTML;

  var btn = document.getElementById("button-add");
  sid = selectedRow.cells[0].innerHTML;
  if (btn) {
      btn.innerHTML = "Update";
      btn.setAttribute("onclick", "update(sid)");
  }
}

function update(sid) {
  var newData = getFormData()
  fetch('/student/'+sid, {
      method: "PUT",
      body: JSON.stringify(newData),
      headers: {"Content-type": "application/json; charset=UTF-8"}
  }).then (res =>{
      if (res.ok) {
          // fill in selected row with updated value
          selectedRow.cells[0].innerHTML = newData.stdid;
          selectedRow.cells[1].innerHTML = newData.fname;
          selectedRow.cells[2].innerHTML = newData.lname;
          selectedRow.cells[3].innerHTML = newData.email;
          // set to previous value
          var button = document.getElementById("button-add");
          button.innerHTML = "Add";
          button.setAttribute("onclick", "addStudent()")
          selectedRow = null;

          resetform();
      }else {
          alert("Server: Update request error.")
      }
  })
}

function deleteStudent(r){
  // this(input) -> td -> tr
  if (confirm('Are you sure you want to DELETE this?')){
      selectedRow = r.parentElement.parentElement;
      sid = selectedRow.cells[0].innerHTML;

      fetch('/student/' +sid,{
          method: "DELETE",
          headers: {"Content-type": "application/json; charset=UTF-8"}
      });
      var rowIndex = selectedRow.rowIndex;    // index starts from 0
      if (rowIndex>0){
          // td is row 0
          document.getElementById("myTable").deleteRow(rowIndex);
      }
      selectedRow = null;
  }
}

function newRow(student) {

  // Find a <table> element with id = "myTable":
  var table = document.getElementById("myTable");
      
  // Create an empty <tr> element and add to the last position of the table:
  var row = table.insertRow(table.length);

  // insert new cells (<td> element) at the 1st and 2nd position of the "new" <tr> element:
  var td = []

  // iterate the loop till the roe 0 cell length in the table
  for(i=0; i<table.rows[0].cells.length; i++){
      td[i] = row.insertCell(i);
  }
      // Add student detail to the new cells: 
  td[0].innerHTML = student.stdid; 
  td[1].innerHTML = student.fname; 
  td[2].innerHTML = student.lname; 
  td[3].innerHTML = student.email; 
  td[4].innerHTML = '<input type="button" onclick="deleteStudent(this)" value="delete" id="button-1">'; 
  td[5].innerHTML = '<input type="button" onclick="updateStudent(this)" value="edit" id="button-2">';
}

function getFormData(){
  var formData = {
      stdid : parseInt(document.getElementById("sid").value),
      fname : document.getElementById("fname").value,
      lname : document.getElementById("lname").value,
      email : document.getElementById("email").value
  }
  return formData
}