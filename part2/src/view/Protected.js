import m from "mithril";

import User from "../model/User";

var Protected = {
    view: function() {
        return m("div", {}, [m("h1", "This content is only for logged in users"), m("div", "Hello " + User.current.username)]);
    }
};

module.exports = Protected;
