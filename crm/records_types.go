package crm

type Account struct {
	Data []struct {
		Owner struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Owner,omitempty"`
		Ownership        interface{} `json:"Ownership,omitempty"`
		Description      string      `json:"Description,omitempty"`
		CurrencySymbol   string      `json:"$currency_symbol,omitempty"`
		AccountType      interface{} `json:"Account_Type,omitempty"`
		Rating           string      `json:"Rating,omitempty"`
		SICCode          int         `json:"SIC_Code,omitempty"`
		ShippingState    string      `json:"Shipping_State,omitempty"`
		Website          string      `json:"Website,omitempty"`
		Employees        int         `json:"Employees,omitempty"`
		LastActivityTime string      `json:"Last_Activity_Time,omitempty"`
		Industry         string      `json:"Industry,omitempty"`
		RecordImage      interface{} `json:"Record_Image,omitempty"`
		ModifiedBy       struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Modified_By,omitempty"`
		AccountSite    interface{} `json:"Account_Site,omitempty"`
		ProcessFlow    bool        `json:"$process_flow,omitempty"`
		ExchangeRate   int         `json:"Exchange_Rate,omitempty"`
		Phone          string      `json:"Phone,omitempty"`
		Currency       string      `json:"Currency,omitempty"`
		BillingCountry string      `json:"Billing_Country,omitempty"`
		AccountName    string      `json:"Account_Name,omitempty"`
		ID             string      `json:"id,omitempty"`
		AccountNumber  string      `json:"Account_Number,omitempty"`
		Approved       bool        `json:"$approved,omitempty"`
		TickerSymbol   interface{} `json:"Ticker_Symbol,omitempty"`
		Approval       struct {
			Delegate bool `json:"delegate,omitempty"`
			Approve  bool `json:"approve,omitempty"`
			Reject   bool `json:"reject,omitempty"`
			Resubmit bool `json:"resubmit,omitempty"`
		} `json:"$approval,omitempty"`
		ModifiedTime    string        `json:"Modified_Time,omitempty"`
		BillingStreet   string        `json:"Billing_Street,omitempty"`
		CreatedTime     string        `json:"Created_Time,omitempty"`
		Editable        bool          `json:"$editable,omitempty"`
		BillingCode     string        `json:"Billing_Code,omitempty"`
		Territories     []string      `json:"Territories,omitempty"`
		ParentAccount   interface{}   `json:"Parent_Account,omitempty"`
		ShippingCity    string        `json:"Shipping_City,omitempty"`
		ShippingCountry string        `json:"Shipping_Country,omitempty"`
		ShippingCode    string        `json:"Shipping_Code,omitempty"`
		BillingCity     string        `json:"Billing_City,omitempty"`
		BillingState    string        `json:"Billing_State,omitempty"`
		Tag             []interface{} `json:"Tag,omitempty"`
		CreatedBy       struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Created_By,omitempty"`
		Fax            string `json:"Fax,omitempty"`
		AnnualRevenue  int    `json:"Annual_Revenue,omitempty"`
		ShippingStreet string `json:"Shipping_Street,omitempty"`
	} `json:"data,omitempty"`
	Info PageInfo `json:"info,omitempty"`
}

type Call struct {
	Data []struct {
		CallDuration string `json:"Call_Duration,omitempty"`
		Owner        struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Owner,omitempty"`
		Description    string `json:"Description,omitempty"`
		CurrencySymbol string `json:"$currency_symbol,omitempty"`
		ModifiedBy     struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Modified_By,omitempty"`
		ProcessFlow bool   `json:"$process_flow,omitempty"`
		CallPurpose string `json:"Call_Purpose,omitempty"`
		ID          string `json:"id,omitempty"`
		CallStatus  string `json:"Call_Status,omitempty"`
		Approved    bool   `json:"$approved,omitempty"`
		WhoID       struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Who_Id,omitempty"`
		Approval struct {
			Delegate bool `json:"delegate,omitempty"`
			Approve  bool `json:"approve,omitempty"`
			Reject   bool `json:"reject,omitempty"`
			Resubmit bool `json:"resubmit,omitempty"`
		} `json:"$approval,omitempty"`
		ModifiedTime  string      `json:"Modified_Time,omitempty"`
		Reminder      interface{} `json:"Reminder,omitempty"`
		CreatedTime   string      `json:"Created_Time,omitempty"`
		CallStartTime string      `json:"Call_Start_Time,omitempty"`
		Billable      bool        `json:"Billable,omitempty"`
		Editable      bool        `json:"$editable,omitempty"`
		Subject       string      `json:"Subject,omitempty"`
		SeModule      string      `json:"$se_module,omitempty"`
		CallType      string      `json:"Call_Type,omitempty"`
		CallResult    interface{} `json:"Call_Result,omitempty"`
		WhatID        struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"What_Id,omitempty"`
		CreatedBy struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Created_By,omitempty"`
		Tag []interface{} `json:"Tag,omitempty"`
	} `json:"data,omitempty"`
	Info PageInfo `json:"info,omitempty"`
}

type Campaign struct {
	Data []struct {
		Owner struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Owner,omitempty"`
		Description    string `json:"Description,omitempty"`
		CurrencySymbol string `json:"$currency_symbol,omitempty"`
		CampaignName   string `json:"Campaign_Name,omitempty"`
		EndDate        Date   `json:"End_Date,omitempty"`
		ModifiedBy     struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Modified_By,omitempty"`
		NumSent          string      `json:"Num_sent,omitempty"`
		ProcessFlow      bool        `json:"$process_flow,omitempty"`
		ExchangeRate     int         `json:"Exchange_Rate,omitempty"`
		ExpectedRevenue  interface{} `json:"Expected_Revenue,omitempty"`
		Currency         string      `json:"Currency,omitempty"`
		ActualCost       int         `json:"Actual_Cost,omitempty"`
		ID               string      `json:"id,omitempty"`
		ExpectedResponse interface{} `json:"Expected_Response,omitempty"`
		StartDate        interface{} `json:"Start_Date,omitempty"`
		Approved         bool        `json:"$approved,omitempty"`
		Status           interface{} `json:"Status,omitempty"`
		Approval         struct {
			Delegate bool `json:"delegate,omitempty"`
			Approve  bool `json:"approve,omitempty"`
			Reject   bool `json:"reject,omitempty"`
			Resubmit bool `json:"resubmit,omitempty"`
		} `json:"$approval,omitempty"`
		ModifiedTime string `json:"Modified_Time,omitempty"`
		CreatedTime  string `json:"Created_Time,omitempty"`
		Editable     bool   `json:"$editable,omitempty"`
		Type         string `json:"Type,omitempty"`
		CreatedBy    struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Created_By,omitempty"`
		Tag          []interface{} `json:"Tag,omitempty"`
		BudgetedCost interface{}   `json:"Budgeted_Cost,omitempty"`
	} `json:"data,omitempty"`
	Info struct {
		PerPage     int  `json:"per_page,omitempty"`
		Count       int  `json:"count,omitempty"`
		Page        int  `json:"page,omitempty"`
		MoreRecords bool `json:"more_records,omitempty"`
	} `json:"info,omitempty"`
}

type Case struct {
}

type Contact struct {
	Data []struct {
		EXT   int64 `json:"EXT,omitempty"`
		Owner struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Owner,omitempty"`
		GCLID        interface{} `json:"GCLID,omitempty"`
		LeadScore    int64       `json:"Lead_Score,omitempty"`
		MailingState string      `json:"Mailing_State,omitempty"`
		OtherCountry string      `json:"Other_Country,omitempty"`
		Department   string      `json:"Department,omitempty"`
		ProcessFlow  bool        `json:"$process_flow,omitempty"`
		Currency     string      `json:"Currency,omitempty"`
		AdNetwork    interface{} `json:"Ad_Network,omitempty"`
		ID           string      `json:"id,omitempty"`
		Approval     struct {
			Delegate bool `json:"delegate,omitempty"`
			Approve  bool `json:"approve,omitempty"`
			Reject   bool `json:"reject,omitempty"`
			Resubmit bool `json:"resubmit,omitempty"`
		} `json:"$approval,omitempty"`
		CostPerClick            float64     `json:"Cost_per_Click,omitempty"`
		FirstVisitedURL         interface{} `json:"First_Visited_URL,omitempty"`
		NegativeTouchPointScore int64       `json:"Negative_Touch_Point_Score,omitempty"`
		CreatedTime             string      `json:"Created_Time,omitempty"`
		NegativeScore           int         `json:"Negative_Score,omitempty"`
		AdClickDate             interface{} `json:"Ad_Click_Date,omitempty"`
		LastVisitedTime         Time        `json:"Last_Visited_Time,omitempty"`
		CreatedBy               struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Created_By,omitempty"`
		TouchPointScore         int         `json:"Touch_Point_Score,omitempty"`
		PositiveScore           int         `json:"Positive_Score,omitempty"`
		Description             string      `json:"Description,omitempty"`
		Ad                      interface{} `json:"Ad,omitempty"`
		NumberOfChats           int64       `json:"Number_Of_Chats,omitempty"`
		SearchPartnerNetwork    interface{} `json:"Search_Partner_Network,omitempty"`
		OtherZip                string      `json:"Other_Zip,omitempty"`
		MailingStreet           string      `json:"Mailing_Street,omitempty"`
		AverageTimeSpentMinutes float64     `json:"Average_Time_Spent_Minutes,omitempty"`
		Salutation              string      `json:"Salutation,omitempty"`
		FullName                string      `json:"Full_Name,omitempty"`
		RecordImage             interface{} `json:"Record_Image,omitempty"`
		SkypeID                 interface{} `json:"Skype_ID,omitempty"`
		AccountName             struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Account_Name,omitempty"`
		EmailOptOut                bool          `json:"Email_Opt_Out,omitempty"`
		Keyword                    interface{}   `json:"Keyword,omitempty"`
		OtherStreet                string        `json:"Other_Street,omitempty"`
		Mobile                     string        `json:"Mobile,omitempty"`
		Territories                []interface{} `json:"Territories,omitempty"`
		AdCampaignName             interface{}   `json:"Ad_Campaign_Name,omitempty"`
		LeadSource                 string        `json:"Lead_Source,omitempty"`
		Tag                        []interface{} `json:"Tag,omitempty"`
		ReasonForConversionFailure interface{}   `json:"Reason_for_Conversion_Failure,omitempty"`
		Email                      string        `json:"Email,omitempty"`
		CurrencySymbol             string        `json:"$currency_symbol,omitempty"`
		VisitorScore               string        `json:"Visitor_Score,omitempty"`
		OtherPhone                 string        `json:"Other_Phone,omitempty"`
		OtherState                 string        `json:"Other_State,omitempty"`
		LastActivityTime           string        `json:"Last_Activity_Time,omitempty"`
		ExchangeRate               int           `json:"Exchange_Rate,omitempty"`
		MailingCountry             string        `json:"Mailing_Country,omitempty"`
		Approved                   bool          `json:"$approved,omitempty"`
		ConversionExportedOn       interface{}   `json:"Conversion_Exported_On,omitempty"`
		ClickType                  interface{}   `json:"Click_Type,omitempty"`
		DaysVisited                int64         `json:"Days_Visited,omitempty"`
		OtherCity                  string        `json:"Other_City,omitempty"`
		Editable                   bool          `json:"$editable,omitempty"`
		AdGroupName                interface{}   `json:"AdGroup_Name,omitempty"`
		PositiveTouchPointScore    int           `json:"Positive_Touch_Point_Score,omitempty"`
		HomePhone                  string        `json:"Home_Phone,omitempty"`
		Score                      int           `json:"Score,omitempty"`
		SecondaryEmail             string        `json:"Secondary_Email,omitempty"`
		VendorName                 struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Vendor_Name,omitempty"`
		MailingZip             string      `json:"Mailing_Zip,omitempty"`
		Twitter                string      `json:"Twitter,omitempty"`
		FirstName              string      `json:"First_Name,omitempty"`
		ConversionExportStatus interface{} `json:"Conversion_Export_Status,omitempty"`
		CostPerConversion      float64     `json:"Cost_per_Conversion,omitempty"`
		ModifiedBy             struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Modified_By,omitempty"`
		Phone            string      `json:"Phone,omitempty"`
		ModifiedTime     string      `json:"Modified_Time,omitempty"`
		MailingCity      string      `json:"Mailing_City,omitempty"`
		DeviceType       interface{} `json:"Device_Type,omitempty"`
		Title            string      `json:"Title,omitempty"`
		FirstVisitedTime Time        `json:"First_Visited_Time,omitempty"`
		LastName         string      `json:"Last_Name,omitempty"`
		Referrer         string      `json:"Referrer,omitempty"`
		Fax              string      `json:"Fax,omitempty"`
	} `json:"data,omitempty"`
	Info PageInfo `json:"info,omitempty"`
}

type Deal struct {
	Data []struct {
		Owner struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Owner,omitempty"`
		GCLID                interface{} `json:"GCLID,omitempty"`
		CurrencySymbol       string      `json:"$currency_symbol,omitempty"`
		LastActivityTime     interface{} `json:"Last_Activity_Time,omitempty"`
		ProcessFlow          bool        `json:"$process_flow,omitempty"`
		DealName             string      `json:"Deal_Name,omitempty"`
		ExchangeRate         int         `json:"Exchange_Rate,omitempty"`
		Currency             string      `json:"Currency,omitempty"`
		AdNetwork            interface{} `json:"Ad_Network,omitempty"`
		Stage                string      `json:"Stage,omitempty"`
		ID                   string      `json:"id,omitempty"`
		Approved             bool        `json:"$approved,omitempty"`
		ConversionExportedOn interface{} `json:"Conversion_Exported_On,omitempty"`
		Approval             struct {
			Delegate bool `json:"delegate,omitempty"`
			Approve  bool `json:"approve,omitempty"`
			Reject   bool `json:"reject,omitempty"`
			Resubmit bool `json:"resubmit,omitempty"`
		} `json:"$approval,omitempty"`
		Territory    []interface{} `json:"Territory,omitempty"`
		CostPerClick int           `json:"Cost_per_Click,omitempty"`
		ClickType    interface{}   `json:"Click_Type,omitempty"`
		CreatedTime  string        `json:"Created_Time,omitempty"`
		Editable     bool          `json:"$editable,omitempty"`
		AdGroupName  interface{}   `json:"AdGroup_Name,omitempty"`
		AdClickDate  interface{}   `json:"Ad_Click_Date,omitempty"`
		CreatedBy    struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Created_By,omitempty"`
		Description            interface{} `json:"Description,omitempty"`
		Ad                     interface{} `json:"Ad,omitempty"`
		CampaignSource         interface{} `json:"Campaign_Source,omitempty"`
		SearchPartnerNetwork   interface{} `json:"Search_Partner_Network,omitempty"`
		ClosingDate            string      `json:"Closing_Date,omitempty"`
		ConversionExportStatus string      `json:"Conversion_Export_Status,omitempty"`
		CostPerConversion      int         `json:"Cost_per_Conversion,omitempty"`
		ModifiedBy             struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Modified_By,omitempty"`
		LeadConversionTime   interface{} `json:"Lead_Conversion_Time,omitempty"`
		OverallSalesDuration int         `json:"Overall_Sales_Duration,omitempty"`
		AccountName          interface{} `json:"Account_Name,omitempty"`
		ModifiedTime         string      `json:"Modified_Time,omitempty"`
		Keyword              interface{} `json:"Keyword,omitempty"`
		Amount               int         `json:"Amount,omitempty"`
		DeviceType           interface{} `json:"Device_Type,omitempty"`
		NextStep             string      `json:"Next_Step,omitempty"`
		Probability          int         `json:"Probability,omitempty"`
		ContactName          struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Contact_Name,omitempty"`
		PredictionScore            int           `json:"Prediction_Score,omitempty"`
		SalesCycleDuration         int           `json:"Sales_Cycle_Duration,omitempty"`
		AmountQuoted               interface{}   `json:"Amount_Quoted,omitempty"`
		AdCampaignName             interface{}   `json:"Ad_Campaign_Name,omitempty"`
		LeadSource                 string        `json:"Lead_Source,omitempty"`
		Tag                        []interface{} `json:"Tag,omitempty"`
		ReasonForConversionFailure interface{}   `json:"Reason_for_Conversion_Failure,omitempty"`
	} `json:"data,omitempty"`
	Info PageInfo `json:"info,omitempty"`
}

type Event struct {
}

type Invoice struct {
}

type Lead struct {
}

type Potential struct {
}

type PriceBook struct {
}

type Product struct {
}

type PurchaseOrder struct {
}

type Quote struct {
}

type SalesOrder struct {
}

type Solution struct {
}

type Task struct {
}

type Vendor struct {
}
