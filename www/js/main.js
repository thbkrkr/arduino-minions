!(function () {

   window.onload = function () {

        var actions = document.getElementsByClassName("action");

        var handleTap = function (event) {
            doThat(this, this.getAttribute("action"));
        };

        var handleHold = function () {
            var elem = this;
            var action = elem.getAttribute("action");
            if (action == "volup") doThat(elem, "volupup");
            else if (action == "voldown") doThat(elem, "voldowndown");
        };
        for (var i = 0; i < actions.length; i++) Hammer(actions[i]).on("tap", handleTap);
        for (i = 0; i < actions.length; i++) Hammer(actions[i]).on("hold", handleHold);

    };

    var robot = "minions-bot";
    var apiHost = "http://localhost:3000";
    
    var apiUrl = apiHost + "/api/robots/" + robot + "/devices/";

    function doThat(elem, action) {
        elem.style.opacity = "0.5";

        microAjax(apiUrl + action + "/commands/Toggle", function (response) {
            window.setTimeout(function () {
                elem.style.opacity = "1";
                elem.style.opacity = "";
            }, 50);
        });
    }

})(document);
