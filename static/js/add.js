function GoBack(){
    window.location.href="/"
}

function Save(){
    let student = {};
    student['name'] = $("#txt_student_name")[0].value;
    student['gender'] = $("#sel_student_gender")[0].value;
    student['age'] = $("#txt_student_age")[0].value;
    student['birthday'] = $("#txt_student_birthday")[0].value;
    student['email'] = $("#txt_student_email")[0].value;
    student['phone'] = $("#txt_student_phone")[0].value;
    student['address'] = $("#txt_student_address")[0].value;
    student['id'] = $.getUrlParam('id');
    $.post("/save", student, result=>{
        window.location.href="/"

    })
}

function GeneralFake(){
    faker.locale = "zh_CN";
    let randomName = faker.name.findName(); // Caitlyn Kerluke
    let randomEmail = faker.internet.email(); // Rusty@arne.info
    let randomPhone = faker.phone.phoneNumber();
    let randomAddress = faker.address.streetAddress();
    let randomGender = faker.random.boolean() === true?1:0;
    let odob = faker.date.past(40, new Date("Sat Sep 20 2020 21:35:02 GMT+0200 (CEST)"));
    let dob = odob.getFullYear() + "-" + ((odob.getMonth()+1)<10?'0'+(odob.getMonth()+1):odob.getMonth()+1) + "-" + (odob.getDate()<10?'0'+odob.getDate():odob.getDate());
    let curYear = new Date().getFullYear();

    $("#txt_student_name")[0].value = randomName;
    $("#txt_student_email")[0].value = randomEmail;
    $("#txt_student_phone")[0].value = randomPhone;
    $("#txt_student_address")[0].value = randomAddress;
    $("#sel_student_gender")[0].value = randomGender;
    $("#txt_student_age")[0].value = curYear - odob.getFullYear();
    $("#txt_student_birthday")[0].value = dob;

}