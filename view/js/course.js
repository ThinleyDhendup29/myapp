window.onload = function() {
    fetch("/courses")
    .then(response => response.text())
    .then(data => showCourses(data))
}

function addRow(course) {
    var table = document.getElementById("myTable")
    var row = table.insertRow(table.length)
    //insert
    var td = [] 
    //to put 4 data, cid, cname, etc so for loop is used
    for (i = 0; i < table.rows[0].cells.length; i++) {
        td[i] = row.insertCell(i)
    }
    td[0].innerHTML = course.cid
    td[1].innerHTML = course.coursename
    td[2].innerHTML = '<input type="button" onclick="deleteCourse(this)" value="delete" id="button-1">'
    td[3].innerHTML = '<input type="button" onclick="updateCourse(this)" value="edit" id="button-2">'
}

function showCourses(data) {
    const courses = JSON.parse(data)
    courses.forEach(c => {
        addRow(c)
    })
}

// convert json body to js
function showCourse(data) {
    const course = JSON.parse(data)
    addRow(course)
}

// reset the set form fields to empty
function resetForm() {
    document.getElementById("cid").value = ""
    document.getElementById("cname").value = ""
}


//helper function
function getFormData() {
    var data = {
        cid : document.getElementById("cid").value,
        coursename : document.getElementById("cname").value,
    }
    return data
}

function addCourse() {
    var data = getFormData()
    var cId = data.cid
    //if cid in not a number
    if(cId=="") {
        alert("Enter valid course ID")
        return
    } else if (data.coursename == "") {
        alert("Course name cannot be empty")
        return
    }

    //call the api using fetch() which has 2 arguments: url, request body
    fetch('/course', {
        method: "POST",
        //convert the data which have js body to json as server doesnt understand js
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response1 => {
        //promise is resolved
        if (response1.ok) {
            fetch("/course/"+cId)
            .then(response2 => response2.text())
            .then(data => showCourse(data))
        } else {
            throw new Error(response1.status)
        }
    }).catch(e => {
        //if cookie is not found
        if (e.message == 303) {
            alert("User not logged in.")
            window.open("index.html", "_self")
        } else if (e.message == 500){
            alert("Server error!")
        }
    })
    resetForm()
}

function deleteCourse(r){
    if (confirm('Are you sure you want to DELETE this?')){
        selectedRow = r.parentElement.parentElement
        cid = selectedRow.cells[0].innerHTML

        fetch('/course/'+cid, {
            method: "DELETE",
            headers: {"Content-type": "application/json"}
        }).then(res => {
            if (res.ok) {
                alert("Course deleted")
                //delete the row with the row data
                var rowIndex = selectedRow.rowIndex; // index starts from 0
                if (rowIndex>0) { //th is row 0
                    document.getElementById("myTable").deleteRow(rowIndex)
                }
                selectedRow = null
            }
        })
    }
}

function update(cid) {
    //extract new data from form using helper function
    var newData = getFormData()
    fetch('/course/'+cid, {
        method: "PUT",
        body: JSON.stringify(newData),
        headers: {"Content-type":"application/json"}
    }).then (res => {
        // if response is ok = success
        if (res.ok) {
            //display updated data 
            selectedRow.cells[0].innerHTML = newData.cid
            selectedRow.cells[1].innerHTML = newData.coursename

            //set the button value back to add
            var btn = document.getElementById("button-add")
            btn.innerHTML = "Add"
            btn.setAttribute("onclick", "addCourse()")
            
            //initial value null
            selectedRow = null

            resetForm()
        } else {
            alert("Server: Update request error.")
        }
    })
}

//global variable = selectedRow
var selectedRow = null
function updateCourse(r) {
    //r=this(input element), parentElement = td, parentElement = tr
    selectedRow = r.parentElement.parentElement

    //display selected row data in the form
    document.getElementById("cid").value = selectedRow.cells[0].innerHTML
    document.getElementById("cname").value = selectedRow.cells[1].innerHTML

    //change button value from add to update
    var btn = document.getElementById("button-add")

    if (btn) {
        btn.innerHTML = "Update"
        //extract cid as url variable
        cid = selectedRow.cells[0].innerHTML
        btn.setAttribute("onclick", "update(cid)")
    }
}