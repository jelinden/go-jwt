import m from "mithril";
import User from "../model/User";
import LoginService from "../service/Login";

var Login = {
    view: function() {
        return m("div", {}, [
            m("h1", "Please login"),
            m(
                "form",
                {
                    onsubmit: function(e) {
                        e.preventDefault();
                        LoginService.post();
                    }
                },
                [
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
                    m("button.button[type=submit]", "Login")
                ]
            ),
            m("div.error", User.current.error)
        ]);
    }
};

module.exports = Login;
