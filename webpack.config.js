// Generated using webpack-cli https://github.com/webpack/webpack-cli

const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const { profile } = require('console');

const isProduction = process.env.NODE_ENV == 'production';


const stylesHandler = isProduction ? MiniCssExtractPlugin.loader : 'style-loader';


//G:\kursach_PP
const config = {
    entry: {
        index: './src/index.js',
        common: './src/scripts/common.js',
        profile: './src/scripts/profile.js',
        project: './src/scripts/project.js'
    },
    output: {
        path: path.resolve(__dirname, 'dist'),
    },
    devServer: {
        open: true,
        host: 'localhost',
    },
    plugins: [
        new CleanWebpackPlugin(),
        new HtmlWebpackPlugin({
            filename: 'index.html',
            template: 'src/pages/index.html',
            minify: false,
            chunks: ['common', 'index'],
        }),
        new HtmlWebpackPlugin({
            filename: 'profile.html',
            template: 'src/pages/profile.html',
            minify: false,
            chunks: ['common', 'profile'],
        }),
        new HtmlWebpackPlugin({
          filename: 'project.html',
          template: 'src/pages/project.html',
          minify: false,
          chunks: ['common', 'project'],
      }),

        // Add your plugins here
        // Learn more about plugins from https://webpack.js.org/configuration/plugins/
    ],
    module: {
        rules: [
            {
                test: /\.css$/i,
                use: [stylesHandler, 'css-loader', 'postcss-loader'],
            },
            {
                test: /\.(eot|svg|ttf|woff|woff2|png|jpg|gif)$/i,
                type: 'asset',
            },
            {
                test: /\.(png|svg|jpg|jpeg|gif)$/i,
                type: 'asset/resource',
            },
            {
                test: /\.(woff|woff2|eot|ttf|otf)$/i,
                type: 'asset/resource',
            },

            // Add your rules for custom modules here
            // Learn more about loaders from https://webpack.js.org/loaders/
        ],
    },
};

module.exports = () => {
    if (isProduction) {
        config.mode = 'production';

        config.plugins.push(new MiniCssExtractPlugin());


    } else {
        config.mode = 'development';
    }
    return config;
};
