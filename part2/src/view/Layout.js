import m from "mithril";

module.exports = {
    view: function(vnode) {
        return m("main.layout", [
            m("nav.menu", [
                m("a[href='/']", { oncreate: m.route.link }, "Index"),
                m("a[href='/login']", { oncreate: m.route.link }, "Login"),
                m("a[href='/signup']", { oncreate: m.route.link }, "Signup")
            ]),

            m("section", vnode.children)
        ]);
    }
};
