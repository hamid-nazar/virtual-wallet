package errors

import "errors"

// General Errors
var (
	ErrCardNotFound              = errors.New("Card does not exist")
	ErrUserNotFound              = errors.New("User does not exist")
	ErrCategoryNotFound          = errors.New("Category does not exist")
	ErrWalletNotFound            = errors.New("Wallet does not exist")
	ErrPaymentInstrumentNotFound = errors.New("Payment instrument does not exist")
	ErrTransactionNotFound       = errors.New("Transaction does not exist")
)

// Transaction Errors
var (
	ErrTransactionUnconfirmedUser          = errors.New("Please verify your account before making a transaction by clicking on the link in the verification email you received.")
	ErrTransactionLargeAmount              = errors.New("In order to complete your transaction, please click on the verification link that we've just sent to your email for security purposes.")
	ErrInsufficientWalletFunds             = errors.New("You have insufficient funds in your wallet to perform this transaction. Please top-up and try again.")
	ErrNoDefaultWallet                     = errors.New("Unable to perform transaction, because the recipient has no default wallet selected.")
	ErrInsufficientCardFunds               = errors.New("Insufficient funds to complete the transaction.")
	ErrCSVFormatMismatch                   = errors.New("Supplied CSV does not match card CSV.")
	ErrInvalidRecipientContactType         = errors.New("Invalid recipient contact type when creating transaction.")
	ErrExpiredVerificationTokenText        = errors.New("The verification link for your transaction has expired. Please start a new transaction and verify it within one hour, if we require security verification via email.")
	ErrExpiredVerificationTokenTitle       = errors.New("Verification link expired.")
	ErrMakingTransactionWithBlockedAccount = errors.New("Your account is currently blocked and you cannot make any transactions. Please contact Virtual Wallet support.")
	ErrTransactionFailed                   = errors.New("Transaction error")
	ErrTransactionHistoryError             = errors.New("Transaction history error")
)

// Duplication Errors
var (
	ErrDuplicateUsername              = errors.New("User with this username already exists.")
	ErrDuplicateEmail                 = errors.New("User with this email already exists.")
	ErrDuplicatePhoneNumber           = errors.New("This phone number is already in use.")
	ErrDuplicateWallet                = errors.New("Wallet with the same name already exists.")
	ErrDuplicateCardNumber            = errors.New("Card with the same number already exists.")
	ErrDuplicateCardName              = errors.New("Card with the same name already exists for this user.")
	ErrDuplicateCategory              = errors.New("You already have a category with this name.")
	ErrDuplicatePaymentInstrumentName = errors.New("User already has payment instrument with the same name.")
)

// Sorting Errors
var (
	ErrInvalidSortingArguments = errors.New("Criterion and sorting direction required. Use '.' as separator.")
	ErrInvalidSortingCriterion = errors.New("Unsupported sorting criterion. Supported criteria: 'amount' and/or 'date'.")
	ErrInvalidSortingDirection = errors.New("Invalid sorting direction. Valid directions: 'asc' or 'desc'.")
)

// Ownership and Access Errors
var (
	ErrUserDoesNotOwnPaymentInstrument = errors.New("You do not own this payment instrument.")
)

// Account and Wallet Errors
var (
	ErrNoTransactionsWithCounterparty = errors.New("You have not sent and/or received any transactions to/from this person.")
	ErrNoTransactionsMatchingCriteria = errors.New("There are no transactions found that match your criteria.")
	ErrNoTransactionsYet              = errors.New("You have not made any transactions yet.")
	ErrEditingCardWithBlockedAccount  = errors.New("Your account is currently blocked and you cannot edit your card. Please contact Virtual Wallet support.")
	ErrDeletingCardWithBlockedAccount = errors.New("Your account is currently blocked and you cannot delete your card. Please contact Virtual Wallet support.")
)

// Image Errors
var (
	ErrInvalidImageFormat = errors.New("Please upload an image in jpeg or png format.")
	ErrDuringImageUpload  = errors.New("Failed to upload image. Please try again later or contact our support if you are still experiencing this issue.")
)

// Filter and Search Errors
var (
	ErrInvalidFilterTypeParameter  = errors.New("Invalid filter type parameter. Please select 'username', 'email' or 'phone'.")
	ErrNoUserMatchesSearchCriteria = errors.New("No user matches the search criteria.")
)

// Admin Errors
var (
	ErrAdminUsersEmptyTitle   = errors.New("There are currently no registered users.")
	ErrAdminIllegalPhoneValue = errors.New("Invalid phone number value. Please use digits only, and no letters.")
	ErrAdminIllegalEmailValue = errors.New("Invalid email address.")
	ErrAdminIllegalValue      = errors.New("Illegal value.")
)

// Error definitions
var (
	ErrUnconfirmedUserTransaction         = errors.New("Please verify your account before making a transaction by clicking on the link in the verification e-mail you received")
	ErrLargeAmountTransaction             = errors.New("In order to complete your transaction, please click on the verification link we've just sent to your e-mail for security purposes")
	ErrCSVMismatch                        = errors.New("Supplied CSV does not match card CSV")
	ErrExpiredVerificationToken           = errors.New("The verification link for your transaction has expired. Please start a new transaction and verify it within one hour, if we require security verification via e-mail")
	ErrBlockedAccountTransaction          = errors.New("Your account is currently blocked and you cannot make any transactions. Please contact Virtual Wallet support to help you further")
	ErrBlockedAccountCardEdit             = errors.New("Your account is currently blocked and you cannot edit your card. Please contact Virtual Wallet support to help you further")
	ErrBlockedAccountCardDelete           = errors.New("Your account is currently blocked and you cannot delete your card. Please contact Virtual Wallet support to help you further")
	ErrNoTransactionsWithThisCounterparty = errors.New("You have not sent and/or received any transactions to/from this person")
	ErrEditingProfileWrongPassword        = errors.New("You have typed incorrect account password")
	ErrLoginFailed                        = errors.New("Invalid username and/or password")
	ErrConstraintViolation                = errors.New("Invalid input")
	ErrInvalidCredentials                 = errors.New("Invalid username and/or password")
	ErrInvalidOperation                   = errors.New("You are not allowed to perform this action")
	ErrAccessDenied                       = errors.New("Access denied to perform this action")
	ErrLinkInvalidOrBroken                = errors.New("The link is invalid or broken")
	ErrSenderHasNoDefaultWallet           = errors.New("You do not have a default wallet. Please create a wallet and choose it as default to send money")
	ErrRecipientHasNoDefaultWallet        = errors.New("Transaction failed, because the recipient has no default wallet")
	ErrSenderAndReceiverWalletsSame       = errors.New("You cannot send money from your default wallet to your default wallet. Please choose a different wallet to send the money from")
)

// Email-related Messages
const (
	EmailSubjectCompleteRegistration = "Complete your registration at VirtualWallet"
	EmailSender                      = "virtualwalletjava@gmail.com"
	EmailMessageConfirmAccount       = "To confirm your account, please click on the following link: http://localhost:8080/confirm-account?token="
	EmailLargeTransactionSubject     = "Confirm transaction in VirtualWallet"
	EmailLargeTransactionMessage     = "To confirm your transaction, please click on the following link: http://localhost:8080/confirm-transaction?token="
)

func HandleError(err error) (errMsg string) {
	if err != nil {
		return err.Error()
	}
	return
}
