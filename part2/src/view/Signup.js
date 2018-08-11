import m from "mithril";
import User from "../model/Signup";
import Login from "../service/Login";

var Signup = {
    view: function() {
        return m("div", {}, [
            m("h1", "Signup"),
            m("form", { onsubmit: sendForm }, [
                m("label.label", "username"),
                m("input.input[type=text][placeholder=username]", {
                    oninput: m.withAttr("value", function(value) {
                        User.current.username = value;
                    })
                }),
                m("label.label", "password"),
                m("input.input[placeholder=password]", {
                    oninput: m.withAttr("value", function(value) {
                        User.current.password = value;
                    }),
                    type: "password"
                }),
                m("button.button[type=submit]", "Signup")
            ]),
            m("div.error", User.current.error)
        ]);
    }
};

var sendForm = function(e) {
    e.preventDefault();
    Login.signup();
};

module.exports = Signup;
