package entity

//User represents users table in database
type User struct {
	ID       uint64  `db:"id" gorm:"primary_key:auto_increment" json:"id"`
	Name     string `db:"name" gorm:"type:varchar(255)" json:"name"`
	Email    string `db:"email" gorm:"uniqueIndex;type:varchar(50)" json:"email"`
	Password string `db:"password" gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
	//Books        *[]Book `json:"books,omitempty"`
	Verified_Otp string `gorm:"type:varchar(255)" json:"verified_otp"`
	Verified     uint   `json:"verified"`
}

// User is an object representing the database table.
// type User struct {
// 	ID              uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
// 	Name            null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
// 	Email           string      `boil:"email" json:"email" toml:"email" yaml:"email"`
// 	Phone           null.String `boil:"phone" json:"phone,omitempty" toml:"phone" yaml:"phone,omitempty"`
// 	Lga             null.String `boil:"lga" json:"lga,omitempty" toml:"lga" yaml:"lga,omitempty"`
// 	Address         null.String `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
// 	State           null.String `boil:"state" json:"state,omitempty" toml:"state" yaml:"state,omitempty"`
// 	VerifiedOtp     null.String `boil:"verified_otp" json:"verified_otp,omitempty" toml:"verified_otp" yaml:"verified_otp,omitempty"`
// 	EmailVerifiedAt null.Time   `boil:"email_verified_at" json:"email_verified_at,omitempty" toml:"email_verified_at" yaml:"email_verified_at,omitempty"`
// 	Verified        null.Int    `boil:"verified" json:"verified,omitempty" toml:"verified" yaml:"verified,omitempty"`
// 	Password        null.String `boil:"password" json:"password,omitempty" toml:"password" yaml:"password,omitempty"`
// 	Dob             null.Time   `boil:"dob" json:"dob,omitempty" toml:"dob" yaml:"dob,omitempty"`
// 	Sex             null.String `boil:"sex" json:"sex,omitempty" toml:"sex" yaml:"sex,omitempty"`
// 	Nationality     null.String `boil:"nationality" json:"nationality,omitempty" toml:"nationality" yaml:"nationality,omitempty"`
// 	IsActive        int         `boil:"is_active" json:"is_active" toml:"is_active" yaml:"is_active"`
// 	Image           null.String `boil:"image" json:"image,omitempty" toml:"image" yaml:"image,omitempty"`
// 	Provider        null.String `boil:"provider" json:"provider,omitempty" toml:"provider" yaml:"provider,omitempty"`
// 	ProviderID      null.Int    `boil:"provider_id" json:"provider_id,omitempty" toml:"provider_id" yaml:"provider_id,omitempty"`
// 	Status          int         `boil:"status" json:"status" toml:"status" yaml:"status"`
// 	RememberToken   null.String `boil:"remember_token" json:"remember_token,omitempty" toml:"remember_token" yaml:"remember_token,omitempty"`
// 	CreatedAt       null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
// 	UpdatedAt       null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
// 	ExpiresAt       null.Time   `boil:"expires_at" json:"expires_at,omitempty" toml:"expires_at" yaml:"expires_at,omitempty"`
// 	GoalOtp         null.String `boil:"goal_otp" json:"goal_otp,omitempty" toml:"goal_otp" yaml:"goal_otp,omitempty"`
// 	GoalContributor null.Int    `boil:"goal_contributor" json:"goal_contributor,omitempty" toml:"goal_contributor" yaml:"goal_contributor,omitempty"`

// 	R *userR `boil:"-" json:"-" toml:"-" yaml:"-"`
// 	L userL  `boil:"-" json:"-" toml:"-" yaml:"-"`
// }
