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


$( "#add-new-recept" ).click(sendForm());
//$ ( "#recept-file" ).click(sendForm());

function sendForm() {
	console.log( "SendForm" );
	var input = $("#new-recept").val();
	if (input != "") {
		console.log( "SendForm - Request sent" );
		var data = $("#recept-textarea").val();
		var url = "http://localhost/saverecept";
		$.ajax({
		  type: "POST",
		  url: url,
		  data: data,
		  success: success,
		  dataType: "text"
		});
	}	
}

function success(){
	alert('done');
}