module.exports = {
  env: {
    es2021: true,
    node: true,
  },
  extends: ["google", "prettier"],
  parser: "@typescript-eslint/parser",
  parserOptions: {
    ecmaVersion: 12,
    sourceType: "module",
  },
  plugins: ["@typescript-eslint"],
  rules: {
    "require-jsdoc": 0,
    quotes: ["error", "single"],
  },
};
