(function(){

	var getMaxHeight = function(selectorString) {
		var maxHeight = 0;
		$(selectorString).each(function(index, element){
			if ($(this).height() > maxHeight) {
				maxHeight = $(this).height();
			}
		});
		return maxHeight;
	}

	$().ready(function(){
		$('#examplesSection h3').height(getMaxHeight('#examplesSection h3'));
		$('#examplesSection p').height(getMaxHeight('#examplesSection p'));
	});
})();