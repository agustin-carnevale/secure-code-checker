import security from "eslint-plugin-security";

/** @type {import("eslint").FlatConfig[]} */
export default [
  {
    files: ["**/*.js"],
    plugins: {
      security,
    },
    rules: {
      "security/detect-eval-with-expression": "error",
      "security/detect-object-injection": "error",
      "security/detect-unsafe-regex": "error",
    },
  },
];
