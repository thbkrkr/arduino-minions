!(function () {

   window.onload = function () {
        var handleTap = function (event) {
            doThat(this, this.getAttribute("action"));
        };
        var handleHold = function () {
            var elem = this;
            var action = elem.getAttribute("action");
            if (action == "volup") doThat(elem, "volupup");
            else if (action == "voldown") doThat(elem, "voldowndown");
        };
        var actions = document.getElementsByClassName("action");
        for (var i = 0; i < actions.length; i++) Hammer(actions[i]).on("tap", handleTap);
        //for (i = 0; i < actions.length; i++) Hammer(actions[i]).on("hold", handleHold);
    };

    api = "/!/";

    function doThat(elem, action) {
        elem.style.opacity = "0.7";

        microAjax(api + action, function (response) {
            window.setTimeout(function () {
                elem.style.opacity = "1";
                elem.style.opacity = "";
            }, 30);
        });
    }

})(document);
