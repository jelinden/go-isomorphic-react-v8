if (typeof selfjs !== 'undefined') {
    selfjs.handleRequest = function(req, res, data) {
        Router.run(routes, req.path, function(Root, state) {
            var html = React.createFactory(Html)({
                markup: <Root params={state.params} data={JSON.parse(data)}/>
            });

            res.write(React.renderToStaticMarkup(html));
        });
    }
}

if (typeof window !== 'undefined') {
    var xmlhttp = new XMLHttpRequest();

    function render() {
        Router.run(routes, Router.HistoryLocation, function(Root, state) {
            if (state.path === "/") {
                getData(function(data) {
                    renderFactory(Root, state, data);
                }, "/api/data");
            }
            if (state.path === "/about") {
                getData(function(data) {
                    renderFactory(Root, state, data);
                }, "/api/anotherpage");
            }

        });
    }

    function renderFactory(Root, state, data) {
        var body = React.createFactory(Root)({
            params: state.params,
            data: data
        });
        React.render(body, document.body);
    }

    function getData(callback, apiEndpoint) {
        xmlhttp.onreadystatechange = function() {
            if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
                var data = JSON.parse(xmlhttp.responseText);
                callback(data);
            }
        }
        xmlhttp.open("GET", apiEndpoint, true);
        xmlhttp.send();
    }

    window.onload = function() {
        render();
    }
}