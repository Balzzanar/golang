$( document ).ready(function() {
    console.log( "ready!" );
});


$( ".recept-file" ).click(function() {
  var rawrecept = JSON.parse($(this).attr('content'));
  rawrecept = JSON.stringify(rawrecept, null, 2);
  $( ".recept-file" ).each(function() {
  	$(this).removeClass('enabled');
  });
  $(this).addClass('enabled')
  $("#recept-textarea").val(rawrecept);
  $("#new-recept").val($(this).text());
});


$( "#add-new-recept" ).click(function () {
	sendForm('new');	
});
//$ ( "#recept-file" ).click(sendForm());

function sendForm(new_recept) {
	console.log( "SendForm" );
	var input = $("#new-recept").val();
	var data = $("#recept-textarea").val();
	if (new_recept == "new" && input != ""){
		data = '{"ID": "'+input+ '"}';
	}
	
		console.log( "SendForm - Request sent" );
		
		var url = "http://sphann.se:8080/saverecept";
		$.ajax({
		  type: "POST",
		  url: url,
		  data: data,
		  success: success,
		  dataType: "text"
		});
	}

function success(){
	alert('done');
}
