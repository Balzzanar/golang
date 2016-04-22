$( document ).ready(function() {
    console.log( "ready!" );
});

$( ".recept-file" ).click(function() {
  var rawrecept = $(this).attr('content');
  $("#recept-textarea").val(rawrecept);
  alert(rawrecept);
});