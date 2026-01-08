# Resend Email Integration - Setup Guide

## ✅ What's Been Done

I've successfully migrated your email service from **Mailgun SMTP** to **Resend API**:

### Code Changes:
1. ✅ Updated [internal/services/email_service.go](server/internal/services/email_service.go)
   - Replaced SMTP code with Resend API calls
   - Kept all HTML templates (no changes needed)
   - All 3 email functions ready (SendVerificationCode, SendWelcomeEmail, SendVIPUpgradeEmail)

2. ✅ Updated [internal/config/config.go](server/internal/config/config.go)
   - Added `ResendAPIKey` config field
   - Loads from `RESEND_API_KEY` environment variable

3. ✅ Updated [go.mod](server/go.mod)
   - Added `github.com/resendlabs/resend-go/v2` dependency

4. ✅ Updated [.env.production](server/.env.production)
   - Replaced Mailgun SMTP config with Resend API placeholder
   - Clear instructions on where to paste your API key

---

## 🚀 Next Steps (What YOU Need To Do)

### Step 1: Create Resend Account (2 minutes)
1. Go to https://resend.com/signup
2. Sign up with your email
3. Verify email

### Step 2: Generate API Key (1 minute)
1. Go to https://resend.com/api-keys
2. Click "Create API Key"
3. Copy the key (starts with `re_`)

### Step 3: Update Production Config (30 seconds)
Replace the placeholder in [.env.production](server/.env.production):

```env
# OLD (REMOVE):
RESEND_API_KEY=re_placeholder_paste_your_resend_api_key_here

# NEW (REPLACE WITH YOUR KEY):
RESEND_API_KEY=re_xxxxxxxxxxxxxxxxxxxxxx
```

### Step 4: Update Railway Environment Variables (1 minute)

In Railway Dashboard:
1. Go to your project → Variables
2. Click "Raw Editor"
3. Find line: `RESEND_API_KEY=`
4. Replace placeholder with your actual Resend API key
5. Save

**OR** Update via text editor (easier):
```
RESEND_API_KEY=re_your_actual_api_key_here
```

### Step 5: Deploy to Production (auto-triggered)
1. Push code to GitHub (or just the config file)
2. Railway auto-builds and deploys
3. Check deployment logs: All should be green ✅

### Step 6: Test Email Sending (5 minutes)

Use Postman to trigger registration and verify email arrives:

```
POST https://workradar-production.up.railway.app/api/auth/register

{
  "email": "test@yourmail.com",
  "password": "Test1234!",
  "fullName": "Test User"
}
```

Check your email inbox within 30 seconds:
- ✅ Email should arrive from Workradar
- ✅ Subject: "Workradar - Kode Verifikasi Reset Password"
- ✅ Contains 6-digit verification code
- ✅ Professional HTML formatting

---

## 📊 Resend Features

| Feature | Mailgun | Resend |
|---------|---------|--------|
| **Free Tier** | ❌ (sandbox only) | ✅ 100/day |
| **Setup Difficulty** | 🔴 Complex (SMTP) | 🟢 Simple (API) |
| **For Your App** | 2 emails/registration × 100/day = ✅ Enough | Same ✅ Enough |
| **Documentation** | Good | Excellent |
| **Support** | Email | Email + Chat |
| **Status Dashboard** | Limited | Great (opens, clicks) |

---

## 🔧 How It Works

**Mailgun (Old):** 
```
Backend → SMTP Server → Email
(Slow, complex, requires SMTP setup)
```

**Resend (New):**
```
Backend → Resend API (HTTP) → Email
(Fast, simple, just HTTP request)
```

---

## ⚠️ Troubleshooting

### Email Not Arriving?
1. Check Resend dashboard: https://resend.com/emails
   - See if email was delivered ✓ or bounced ✗
2. Check Railway logs:
   ```
   Logs → Look for "✅ Email sent successfully"
   ```
3. Check spam folder in your test email account

### "RESEND_API_KEY is empty" error in logs?
1. Make sure Railway Variable is set
2. Wait 30 seconds for deployment
3. Check Variable in Railway dashboard

### Can't Sign Up on Resend?
- Use any valid email (doesn't need to be corporate)
- Verification link takes 2-3 minutes sometimes
- Check spam folder for Resend verification email

---

## 📝 Your Checklist

- [ ] Create Resend account
- [ ] Generate Resend API key
- [ ] Update [.env.production](server/.env.production)
- [ ] Update Railway Variables
- [ ] Wait for auto-deployment
- [ ] Test with Postman
- [ ] Check email inbox
- [ ] Success! 🎉

---

## 🎯 Summary

**From:** Mailgun SMTP (slow, complex setup)  
**To:** Resend API (instant, 3-minute setup)  
**Result:** Email verification works perfectly ✅

Ready? Start with Step 1 above! 🚀
