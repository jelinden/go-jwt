import m from "mithril";

var NotProtected = {
    view: function() {
        return m("div", {}, [m("h1", "Hello"), m("div", "Content for not logged in users.")]);
    }
};

module.exports = NotProtected;
