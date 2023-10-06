package user

import "gorm.io/gorm"

func filterByEmail(email *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if email == nil || *email == "" {
			return db
		} else {
			return db.Where("email like ?", "%"+*email+"%")
		}
	}
}

func joinBankAccount() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Joins("JOIN game.user_bank_accounts ON user_bank_accounts.user_ksuid = users.ksuid")
	}
}

func joinWallet() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Joins("JOIN game.wallets ON wallets.user_ksuid = users.ksuid")
	}
}

func filterByBankAccountName(bankAccountName *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if bankAccountName == nil || *bankAccountName == "" {
			return db
		} else {
			return db.
				Where("user_bank_accounts.bank_account_name like ?", "%"+*bankAccountName+"%")
		}
	}
}

func filterByWalletCurrentAmount(currentAmount *float64, operator string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if currentAmount == nil {
			return db
		} else {
			return db.
				Where("wallets.current_amount "+operator+" ?", *currentAmount)
		}
	}
}

func filterByBankName(bankName *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if bankName == nil || *bankName == "" {
			return db
		} else {
			return db.
				Where("user_bank_accounts.bank_name like ?", "%"+*bankName+"%")
		}
	}
}

func filterByBankAccountNumber(bankAccountNumber *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if bankAccountNumber == nil || *bankAccountNumber == "" {
			return db
		} else {
			return db.
				Where("user_bank_accounts.bank_account_name like ?", "%"+*bankAccountNumber+"%")
		}
	}
}

func filterInternalCompanyPic(employeeKsuid *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if employeeKsuid == nil || *employeeKsuid == "" {
			return db
		} else {
			return db.
				Joins("JOIN company.company_partners ON companies.ksuid = company_partners.partner_ksuid").
				Joins("JOIN company.company_pics ON company_pics.company_partner_ksuid = company_partners.ksuid").
				Where("company_pics.employee_ksuid = ?", *employeeKsuid)
		}
	}
}
