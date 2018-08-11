const UglifyJsPlugin = require("uglifyjs-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const OptimizeCSSAssetsPlugin = require("optimize-css-assets-webpack-plugin");

module.exports = {
    mode: "production",
    context: __dirname,
    entry: { main: "./src/index.js" },
    output: {
        path: __dirname + "/bin",
        filename: "app.js"
    },
    module: {
        rules: [
            {
                test: /\.js/,
                loader: "babel-loader",
                exclude: /node_modules/,
                include: __dirname + "/src"
            },
            {
                test: /\.css$/,
                use: [MiniCssExtractPlugin.loader, "css-loader"]
            }
        ]
    },
    optimization: {
        minimizer: [
            new UglifyJsPlugin({
                cache: true,
                parallel: true,
                sourceMap: false
            }),
            new OptimizeCSSAssetsPlugin({})
        ]
    },
    plugins: [
        new MiniCssExtractPlugin({
            filename: "app.css"
        })
    ]
};
