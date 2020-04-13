function GotoAddPage(){
    window.location.href="/add";
}

function Search(){
    let searchkey = $("#txtSearchKey")[0].value;
    if (searchkey.trim()===""){
        window.location.href="/"
    } else{
        window.location.href="/search?searchkey="+searchkey
    }

}

function Delete(id){
    $.ajax({
        url:"/delete?id="+id,
        type: "delete",
        success: result=>{
            window.location.href="/";
        }
    })
}

function ConfirmDelete(id) {
    $('.ui.basic.modal')
        .modal({
            closable: false,
            onApprove: function () {
                Delete(id)
            }
        })
        .modal('show')
    ;
}

function GotoDetail(id) {
    window.location.href="/detail?id="+id
}


function GoNextPage(page){
    let searchKey =$.getUrlParam('searchkey');
    if (searchKey===null){
        window.location.href="/?page="+page
    } else{
        window.location.href="/search?searchkey="+searchKey+"&page="+page
    }

}