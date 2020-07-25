export class AppError extends Error {
  constructor(message = null) {
    super(message || 'Application Error')
  }
}

export class SignInError extends AppError {
  constructor(error, message = null) {
    super(message || 'Faild to signin')
  }
}
