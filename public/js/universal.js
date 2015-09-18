var React = require('react');
var Router = require('react-router');

var Route = Router.Route;
var DefaultRoute = Router.DefaultRoute;
var RouteHandler = Router.RouteHandler;
var Link = Router.Link;

var Index = React.createClass({
    render: function() {
        return (
            <div>
                <h2>{this.props.data.Title}</h2>
                <NewsList data={this.props.data} />
            </div>
        );
    }
});

var NewsList = React.createClass({
    render: function () {
        return (
            <div>
                {
                    this.props.data.ItemList.map(function (item, i) { 
                        return ( 
                            <NewsItem title={item.Title}
                                pubDate={item.PubDate}
                                link={item.Link} key={i}/> 
                        ); 
                    })
                }
            </div>
        );
    }
});

var NewsItem = React.createClass({
    render: function () {
        return (
            <div className="newsItem">
                <div className="pubDate">{this.props.pubDate}</div>
                <a href={this.props.link}><div className="title">{this.props.title}</div></a>
            </div>
        );
    }
});

var About = React.createClass({
    render: function() {
        return (
            <div>
                <h2>Another page</h2>
                <div>{this.props.data.text}</div>
            </div>
        );
    }
});

var Layout = React.createClass({
    render () {
        return (
            <div>
                <div className="pure-menu pure-menu-horizontal">
                    <Link to="/" className="pure-menu-heading pure-menu-link">Home</Link>
                    <Link to="/about" className="pure-menu-heading pure-menu-link">Another page</Link>
                </div>
                <div className="main">
                    <h1>GO - REACT - ISOMORPHIC</h1>
                    <RouteHandler data={this.props.data}/>
                </div>
            </div>
        );
    }
});

var Html = React.createClass({
    render: function() {
        return (
            <html>
                <head>
                    <meta charSet="utf-8" />
                    <meta httpEquiv="x-ua-compatible" content="ie=edge" />
                    <title>Go - React - Isomorphic</title>
                    <link rel="shortcut icon" href="favicon.ico" />
                    <meta name="description" content="" />
                    <meta name="viewport" content="width=device-width, initial-scale=1" />
                    <link rel="stylesheet" href="/public/css/pure-min.css" />
                    <link rel="stylesheet" href="/public/css/index.css" />
                </head>
                <body>
                    {this.props.markup}
                    <script src="/universal.js" async></script>
                </body>
             </html>
        );
    }
});

var routes = (
    <Route handler={Layout} path="/">
        <DefaultRoute handler={Index}/>
        <Route path="about" handler={About}/>
    </Route>
);

module.exports = {
    Html: Html,
    routes: routes
};