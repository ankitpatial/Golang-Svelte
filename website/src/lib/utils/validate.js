/**
 * validatePwd string
 * @param value {string}
 * @returns {null|string}
 */
export const validatePwd = (value) => {
  const isNonWhiteSpace = /^\S*$/;
  if (!isNonWhiteSpace.test(value)) {
    return "Password must not contain Whitespaces.";
  }

  const isContainsUppercase = /^(?=.*[A-Z]).*$/;
  if (!isContainsUppercase.test(value)) {
    return "Password must have at least one Uppercase Character.";
  }

  const isContainsLowercase = /^(?=.*[a-z]).*$/;
  if (!isContainsLowercase.test(value)) {
    return "Password must have at least one Lowercase Character.";
  }

  const isContainsNumber = /^(?=.*[0-9]).*$/;
  if (!isContainsNumber.test(value)) {
    return "Password must contain at least one Digit.";
  }

  const isContainsSymbol =
    /^(?=.*[~`!@#$%^&*()--+={}\[\]|\\:;"'<>,.?/_₹]).*$/;
  if (!isContainsSymbol.test(value)) {
    return "Password must contain at least one Special Symbol.";
  }

  const isValidLength = /^.{8,50}$/;
  if (!isValidLength.test(value)) {
    return "Password must be 8-50 Characters Long.";
  }

  return null;
}


/**
 * extract yup schema errors
 * @param err {any}
 * @returns {*}
 */
export function extractErrors(err) {
  return err.inner.reduce((acc, err) => {
    return {...acc, [err.path]: err.message};
  }, {});
}
