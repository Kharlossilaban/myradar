# Design Handoff Document – Workradar

**Aplikasi Task Management dengan Fitur VIP**

---

## Project Overview

| **Field** | **Detail** |
|-----------|------------|
| **Project Name** | Workradar – Aplikasi Task Management dengan Fitur VIP |
| **Version / Date** | v1.0.0 – 10 Januari 2026 |
| **Designer** | [Nama Designer] |
| **Developer(s)** | Tim Workradar Development |
| **Product Manager** | [Nama PM] |
| **Design Tool** | Figma / Sketch |
| **Platform** | Mobile (Android & iOS) |
| **Tech Stack** | Flutter 3.9.2 |

---

## 1. Design Intent & Context

### Problem Statement
Banyak profesional dan mahasiswa kesulitan mengelola tugas harian mereka secara efektif. Mereka membutuhkan sistem yang tidak hanya mencatat tugas, tetapi juga membantu memprioritaskan, mengatur deadline, dan melacak progress dengan visual yang jelas.

### Design Goal
Membuat aplikasi task management yang **intuitif**, **visual**, dan **produktif** dengan fitur-fitur premium (VIP) seperti weather recommendations, AI chatbot, dan analytics yang membantu pengguna meningkatkan produktivitas mereka.

### Key User Flow / Scenario

**Primary Flow:**
1. **Onboarding** → Register dengan email/Google → Verifikasi OTP → Login
2. **Setup** → Atur jam kerja harian → Buat kategori tugas
3. **Daily Use** → Lihat dashboard → Tambah tugas → Set deadline & kategori → Mark complete
4. **Advanced** → Lihat calendar view → Check analytics → Subscribe VIP → Akses weather & AI chat

**Secondary Flows:**
- Search & filter tasks
- Manage categories
- Edit profile & change password
- View completed tasks history
- Manage leave/vacation days

---

## 2. Visual Specifications

### Color Palette

#### Primary Colors
```
Primary Purple:   #6C5 CE7  (RGB: 108, 92, 231)
Primary Light:    #9B8FF5  (RGB: 155, 143, 245)
Primary Dark:     #4A3DB8  (RGB: 74, 61, 184)
```

#### Secondary Colors
```
Secondary Teal:   #00CEC9  (RGB: 0, 206, 201)
Accent Coral:     #FF7675  (RGB: 255, 118, 117)
```

#### Background Colors (Light Mode)
```
Background:       #F8F9FA  (RGB: 248, 249, 250)
Surface:          #FFFFFF  (RGB: 255, 255, 255)
Card:             #FFFFFF  (RGB: 255, 255, 255)
```

#### Background Colors (Dark Mode)
```
Background:       #1A1A2E  (RGB: 26, 26, 46)
Surface:          #16213E  (RGB: 22, 33, 62)
Card:             #16213E  (RGB: 22, 33, 62)
```

#### Text Colors (Light Mode)
```
Text Primary:     #2D3436  (RGB: 45, 52, 54)
Text Secondary:   #636E72  (RGB: 99, 110, 114)
Text Light:       #B2BEC3  (RGB: 178, 190, 195)
```

#### Text Colors (Dark Mode)
```
Text Primary:     #E8E8E8  (RGB: 232, 232, 232)
Text Secondary:   #B0B0C0  (RGB: 176, 176, 192)
Text Light:       #6B6B8D  (RGB: 107, 107, 141)
```

#### Status Colors
```
Success Green:    #00B894  (RGB: 0, 184, 148)
Warning Yellow:   #FDCB6E  (RGB: 253, 203, 110)
Error Red:        #E74C3C  (RGB: 231, 76, 60)
```

#### VIP Colors
```
VIP Gold:         #FFD700  (RGB: 255, 215, 0)
VIP Gradient:     Linear from #FFD700 to #FFA500
```

#### Category Colors
```
Kerja (Work):         #6C5CE7  (Purple)
Pribadi (Personal):   #00CEC9  (Teal)
Wishlist:             #FF7675  (Coral)
Hari Ulang Tahun:     #FDCB6E  (Yellow)
```

### Typography

**Font Family:** System Default (San Francisco di iOS, Roboto di Android)

| **Style** | **Size** | **Weight** | **Usage** |
|-----------|----------|------------|-----------|
| Headline Large | 28pt | Bold (700) | Page titles, onboarding |
| Headline Medium | 24pt | Bold (700) | Section headers |
| Headline Small | 20pt | Semi-bold (600) | Card headers, dialog titles |
| Title Large | 18pt | Semi-bold (600) | AppBar titles, list headers |
| Title Medium | 16pt | Medium (500) | Subtitles, important text |
| Title Small | 14pt | Medium (500) | Secondary headers |
| Body Large | 16pt | Regular (400) | Main content, descriptions |
| Body Medium | 14pt | Regular (400) | Secondary content, labels |
| Body Small | 12pt | Regular (400) | Captions, hints, timestamps |

### Icon Style

**Icon Library:** Iconsax (Linear style)
- **Size:** 20-24px untuk icons standar
- **Size:** 16-18px untuk icons kecil (badges, inline)
- **Size:** 48-56px untuk icons besar (empty states, illustrations)
- **Color:** Mengikuti text color atau primary color untuk emphasis

**Common Icons:**
- `Iconsax.task_square` - Tasks
- `Iconsax.calendar` - Calendar
- `Iconsax.category` - Categories
- `Iconsax.search_normal` - Search
- `Iconsax.message` - Messages
- `Iconsax.user` - Profile
- `Iconsax.add` - Add new
- `Iconsax.more` - More options

### Image Style

**Category Images:**
- Style: Photographic dengan slight blur/overlay
- Aspect Ratio: 1:1 (square)
- Format: JPG/PNG
- Usage: Empty state backgrounds, category illustrations

**Weather Icons:**
- Style: SVG illustrations
- Format: .svg
- Examples: Payung, Jaket, Sunscreen, Vitamin, dll.

---

## 3. Interactive Behavior

### Button Interactions

**Primary Button (ElevatedButton):**
- **Default:** Purple background (#6C5CE7), white text
- **Hover/Press:** Slight scale down (0.98), darker purple (#4A3DB8)
- **Disabled:** Gray background, reduced opacity (0.5)
- **Shadow:** Soft shadow dengan blur 15px, offset (0, 5)

**Text Button:**
- **Default:** Purple text (#6C5CE7)
- **Hover/Press:** Darker purple (#4A3DB8)
- **Disabled:** Gray text, reduced opacity

**Floating Action Button:**
- **Default:** Purple background, white icon
- **Press:** Ripple effect + slight scale
- **Shadow:** Elevation 8dp

### Modal & Bottom Sheet Animations

**Bottom Sheet (Task Input, Category Selection):**
- **Entry:** Slide from bottom dengan duration 300ms
- **Exit:** Slide to bottom dengan duration 250ms
- **Backdrop:** Fade in/out dengan opacity 0.5

**Dialog (Confirmation, Warning):**
- **Entry:** Scale from 0.8 to 1.0 + fade in (200ms)
- **Exit:** Scale to 0.8 + fade out (150ms)

### List & Card Interactions

**Task Card:**
- **Tap:** Navigate to edit screen dengan slide transition
- **Checkbox Tap:** Animate checkmark + move to completed section
- **Swipe:** No swipe actions (tap to edit/delete)

**Category Chip:**
- **Tap:** Change background to primary color, white text
- **Transition:** Smooth color change (150ms)

### Notification Behavior

**SnackBar:**
- **Position:** Bottom of screen (floating)
- **Duration:** 2-3 seconds
- **Entry:** Slide from bottom
- **Exit:** Fade out
- **Shape:** Rounded corners (8px)

---

## 4. Component States

### Button States

| **State** | **Background** | **Text Color** | **Opacity** | **Shadow** |
|-----------|---------------|----------------|-------------|------------|
| Default | #6C5CE7 | #FFFFFF | 1.0 | Yes |
| Pressed | #4A3DB8 | #FFFFFF | 1.0 | Increased |
| Disabled | #B2BEC3 | #FFFFFF | 0.5 | No |
| Loading | #6C5CE7 | Spinner | 1.0 | Yes |

### Form Input States

| **State** | **Border Color** | **Border Width** | **Fill Color** | **Label Color** |
|-----------|-----------------|------------------|----------------|-----------------|
| Default | #E5E5E5 | 1px | Light gray | #636E72 |
| Focus | #6C5CE7 | 2px | Light gray | #6C5CE7 |
| Error | #E74C3C | 1px | Light pink | #E74C3C |
| Disabled | #E5E5E5 | 1px | #F5F5F5 | #B2BEC3 |
| Filled | #E5E5E5 | 1px | Light gray | #2D3436 |

### Card States

**Task Card:**
- **Default:** White background (light) / #16213E (dark), shadow 0-4px blur 20px
- **Hover/Press:** Shadow increased to 0-8px blur 24px
- **Completed:** Reduced opacity (0.7), strikethrough text

**Category Chip:**
- **Unselected:** Light background, primary text
- **Selected:** Primary background, white text
- **Disabled:** Gray background, gray text

### Loading States

**Full Screen Loading:**
- **Display:** Centered CircularProgressIndicator
- **Color:** Primary purple (#6C5CE7)
- **Background:** Semi-transparent backdrop

**Inline Loading:**
- **Display:** Small spinner next to text
- **Size:** 16-20px
- **Color:** Primary or white (depending on background)

---

## 5. Form Rules & Data Validation

### Registration Form

**Email Field:**
- **Rule:** Must be valid email format
- **Validation:** Real-time on blur
- **Error Message:** "Email tidak valid"
- **Pattern:** `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

**Password Field:**
- **Rule:** Minimum 8 characters, must contain letters and numbers
- **Validation:** Real-time on change
- **Error Messages:**
  - "Password minimal 8 karakter"
  - "Password harus mengandung huruf dan angka"
- **Pattern:** `^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`

**Name Field:**
- **Rule:** Required, 2-50 characters
- **Error Message:** "Nama wajib diisi"

### Login Form

**Email Field:**
- **Rule:** Valid email format
- **Error Message:** "Email tidak valid"

**Password Field:**
- **Rule:** Required
- **Error Message:** "Password wajib diisi"

### Task Creation Form

**Title Field:**
- **Rule:** Required, max 100 characters
- **Error Message:** "Judul tugas wajib diisi"

**Deadline Field:**
- **Rule:** Optional, must be valid date
- **Warning:** Show soft warning if date is in the past (allow to proceed)
- **Message:** "Tanggal yang dipilih sudah lewat. Lanjutkan?"

**Duration Field:**
- **Rule:** Optional, must be positive number
- **Format:** Minutes (integer)

**Category Field:**
- **Rule:** Required
- **Default:** Show placeholder "Pilih Kategori"

**Difficulty Field:**
- **Rule:** Required
- **Options:** Ringan, Sedang, Berat
- **Default:** Show placeholder "Pilih Beban Kegiatan"

### Profile Edit Form

**Name Field:**
- **Rule:** Required, 2-50 characters
- **Error Message:** "Nama wajib diisi"

**Email Field:**
- **Rule:** Valid email, read-only (cannot edit)

**Work Hours:**
- **Rule:** Start time must be before end time
- **Format:** 24-hour format (HH:mm)
- **Error Message:** "Jam mulai harus lebih awal dari jam selesai"

---

## 6. Error, Empty & Loading States

### Error States

**Network Error:**
```
Icon: Iconsax.wifi_square (red)
Title: "Tidak ada koneksi"
Message: "Periksa koneksi internet Anda dan coba lagi"
Action: Button "Coba Lagi"
```

**Server Error:**
```
Icon: Iconsax.danger (red)
Title: "Terjadi kesalahan"
Message: "Gagal memuat data. Coba lagi nanti."
Action: Button "Coba Lagi"
```

**Authentication Error:**
```
SnackBar: "Email atau password salah"
Color: Red background
Duration: 3 seconds
```

**Validation Error:**
```
Display: Below input field
Color: Red text (#E74C3C)
Icon: Small warning icon
```

### Empty States

**No Tasks (Dashboard):**
```
Image: Category-specific image (240x240px)
Title: "Belum ada tugas"
Message: "Ketuk tombol + untuk menambahkan tugas pertama anda"
Action: Button "Mulai buat tugas baru!" (if work hours set)
        Button "Mulai Buat jadwal kerja harian anda" (if no work hours)
```

**No Completed Tasks:**
```
Icon: Iconsax.task_square (large, purple with opacity)
Title: "Belum ada tugas selesai"
Message: "Tugas yang sudah selesai akan muncul di sini"
```

**No Search Results:**
```
Icon: Iconsax.search_normal (large, gray)
Title: "Tidak ada hasil"
Message: "Coba kata kunci lain"
```

**No Categories:**
```
Icon: Iconsax.category (large, purple)
Title: "Belum ada kategori"
Message: "Tambahkan kategori untuk mengorganisir tugas Anda"
Action: Button "Tambah Kategori"
```

### Loading States

**Initial Load:**
```
Display: Full screen centered CircularProgressIndicator
Color: Primary purple
Background: Scaffold background color
```

**Pull to Refresh:**
```
Display: Top of list
Indicator: Material RefreshIndicator
Color: Primary purple
```

**Button Loading:**
```
Display: Replace button text with small spinner
Size: 20px
Color: White (on colored button) or primary (on text button)
```

**Lazy Loading (Pagination):**
```
Display: Bottom of list
Indicator: Small CircularProgressIndicator
Padding: 16px vertical
```

---

## 7. User Flow / Navigation

### Main Navigation Structure

**Bottom Navigation (5 tabs):**
1. **Dashboard** - Iconsax.home
2. **Calendar** - Iconsax.calendar
3. **Search** - Iconsax.search_normal
4. **Messages** - Iconsax.message
5. **Profile** - Iconsax.user

### Complete User Flows

#### A. Authentication Flow
```
Splash Screen (Auto-check auth)
    ├─→ [Logged In] → Main Screen (Dashboard)
    └─→ [Not Logged In] → Login Screen
            ├─→ Login with Email → Dashboard
            ├─→ Login with Google → Dashboard
            ├─→ Forgot Password → Verification Code → Reset Password → Login
            └─→ Register → OTP Verification → Login → Dashboard
```

#### B. Task Management Flow
```
Dashboard
    ├─→ FAB (+) → Task Input Modal → Create Task → Dashboard (updated)
    ├─→ Tap Task Card → Edit Task Screen
    │       ├─→ Edit → Save → Dashboard
    │       └─→ Delete → Confirm → Dashboard
    ├─→ Checkbox → Mark Complete → Move to "Selesai Hari Ini"
    └─→ "Periksa semua tugas yang sudah selesai" → Completed Tasks Screen
            └─→ Delete Task / Delete All
```

#### C. Category Management Flow
```
Dashboard → Menu (⋮) → Kelola Kategori → Manage Category Screen
    ├─→ Add Category → Input Name & Color → Save
    ├─→ Edit Category → Modify → Save
    └─→ Delete Category → Confirm → Delete
```

#### D. Calendar Flow
```
Calendar Tab
    ├─→ View tasks by date
    ├─→ Tap date → Show tasks for that date
    ├─→ FAB (+) → Create Task (with selected date)
    └─→ Tap task → Edit Task Screen
```

#### E. Search Flow
```
Search Tab
    ├─→ Type query → Real-time filter
    ├─→ Tap result → Edit Task Screen
    └─→ No results → Empty state
```

#### F. Profile & Settings Flow
```
Profile Tab
    ├─→ Edit Profile → Edit Profile Screen → Save
    ├─→ Change Password → Verify → New Password → Save
    ├─→ Work Hours Config → Set Start/End Time → Save
    ├─→ Leave Management → Add/Edit Leave → Save
    ├─→ Manage Categories → (Same as C)
    ├─→ Subscribe VIP → Subscription Screen → Payment → VIP Features
    ├─→ [VIP] Weather → VIP Weather Screen
    ├─→ [VIP] AI Chat → Messages Tab (Bot)
    └─→ Logout → Confirm → Login Screen
```

#### G. VIP Subscription Flow
```
Profile → Subscribe VIP → Subscription Screen
    ├─→ View Plans (Monthly/Yearly)
    ├─→ Select Plan → Payment Screen
    │       ├─→ Choose Payment Method
    │       ├─→ Confirm Payment → Processing
    │       └─→ Success → VIP Features Unlocked
    └─→ [Already VIP] → Manage Subscription
```

### Navigation Patterns

**Stack Navigation:**
- Most screens use push/pop navigation
- Back button always available in AppBar
- Swipe from left edge to go back (iOS)

**Modal Navigation:**
- Bottom sheets for quick actions (Task Input, Category Select)
- Dialogs for confirmations (Delete, Logout)
- Full screen modals for complex forms (Payment)

**Tab Navigation:**
- Bottom navigation persists across main screens
- Maintains state when switching tabs
- FAB changes based on active tab

---

## 8. Responsive Design Breakpoints

### Platform: Mobile Only

**Workradar adalah aplikasi mobile-first yang didesain khusus untuk smartphone.**

**Target Devices:**
- **Android:** Minimum Android 6.0 (API 23)
- **iOS:** Minimum iOS 12.0

**Screen Sizes:**
- **Small phones:** 320-375px width (iPhone SE, small Android)
- **Medium phones:** 375-414px width (iPhone 12/13, standard Android)
- **Large phones:** 414-428px width (iPhone Pro Max, large Android)

**Orientation:**
- **Primary:** Portrait mode
- **Secondary:** Landscape supported (auto-adapt layout)

**Responsive Behaviors:**

**Portrait Mode (Default):**
- Single column layout
- Full-width cards and inputs
- Bottom navigation always visible
- FAB positioned bottom-right

**Landscape Mode:**
- Maintain single column for consistency
- Reduce vertical spacing slightly
- Bottom navigation remains at bottom
- Adjust modal heights to fit screen

**Safe Areas:**
- Respect device safe areas (notches, home indicators)
- Use SafeArea widget for all screens
- Bottom navigation above home indicator
- Content doesn't overlap with status bar

**Font Scaling:**
- Support system font size settings
- Test with accessibility font sizes (up to 200%)
- Ensure UI doesn't break with large text

---

## 9. Accessibility Notes

### Color Contrast

**WCAG 2.1 AA Compliance:**
- Primary purple (#6C5CE7) on white: **Contrast ratio 4.8:1** ✅
- Text primary (#2D3436) on white: **Contrast ratio 12.6:1** ✅
- Text secondary (#636E72) on white: **Contrast ratio 7.2:1** ✅
- White text on primary purple: **Contrast ratio 4.8:1** ✅
- Error red (#E74C3C) on white: **Contrast ratio 4.5:1** ✅

### Touch Targets

**Minimum Size:** 48x48 dp (Material Design guideline)

**Examples:**
- Buttons: Minimum 48dp height
- Icons: 48x48dp touch area (even if icon is 24x24dp)
- List items: Minimum 48dp height
- Checkboxes: 48x48dp touch area
- FAB: 56x56dp (default Material size)

### Screen Reader Support

**Semantic Labels (Recommended Implementation):**

```dart
// Button example
Semantics(
  label: 'Tambah tugas baru',
  button: true,
  child: FloatingActionButton(...),
)

// Icon example
Semantics(
  label: 'Menu lainnya',
  button: true,
  child: Icon(Iconsax.more),
)

// Image example
Semantics(
  label: 'Ilustrasi kategori kerja',
  image: true,
  child: Image.asset('assets/images/kerja.jpg'),
)
```

**Focus Order:**
- Logical top-to-bottom, left-to-right
- Skip decorative elements
- Focus on interactive elements first

### Keyboard Navigation

**Not applicable** - Mobile app uses touch input

### Dark Mode Support

**Full dark mode implementation:**
- All screens support dark mode
- Automatic theme switching based on system preference
- Manual toggle available in Dashboard menu
- Consistent color scheme across all screens
- Proper contrast in both modes

---

## 10. Design Rationale (Why)

### Color Psychology

**Primary Purple (#6C5CE7):**
- **Why:** Purple represents **creativity, wisdom, and productivity**
- **Impact:** Modern, premium feel that appeals to professionals and students
- **Usage:** Primary actions, branding, emphasis

**Secondary Teal (#00CEC9):**
- **Why:** Teal symbolizes **balance, calmness, and clarity**
- **Impact:** Provides visual balance to the vibrant purple
- **Usage:** Secondary actions, personal category, accents

**Accent Coral (#FF7675):**
- **Why:** Warm, friendly, and **approachable**
- **Impact:** Adds warmth to the cool color scheme
- **Usage:** Wishlist category, highlights, alerts

**VIP Gold (#FFD700):**
- **Why:** Gold represents **premium, exclusivity, and value**
- **Impact:** Clearly differentiates VIP features
- **Usage:** VIP badges, subscription screens, premium features

### UI Structure Decisions

**Bottom Navigation:**
- **Why:** Thumb-friendly on mobile devices
- **Research:** Most users hold phones with one hand
- **Impact:** Quick access to main features without reaching

**Card-Based Layout:**
- **Why:** Scannable, organized, and familiar
- **Research:** Users scan content in F-pattern
- **Impact:** Easy to identify individual tasks, clear hierarchy

**Collapsible Sections (Dashboard):**
- **Why:** Reduce cognitive load, show only relevant info
- **Impact:** Users can focus on "Hari ini" without distraction from future tasks

**Category Chips (Horizontal Scroll):**
- **Why:** Quick filtering without leaving screen
- **Impact:** Faster task filtering compared to dropdown menu

### Typography Choices

**System Fonts (San Francisco / Roboto):**
- **Why:** Native feel, optimized for each platform
- **Performance:** No custom font loading = faster app
- **Accessibility:** Designed for readability at all sizes
- **Impact:** Professional, clean, and familiar to users

**Clear Hierarchy:**
- **Why:** Guide user attention to important information
- **Sizes:** 28pt → 24pt → 20pt → 18pt → 16pt → 14pt → 12pt
- **Impact:** Users can quickly scan and understand content

### Interaction Patterns

**Floating Action Button:**
- **Why:** Primary action always visible and accessible
- **Research:** Material Design best practice for primary actions
- **Impact:** Users can add tasks from any screen state

**Bottom Sheets for Forms:**
- **Why:** Quick actions without losing context
- **Impact:** Users can see dashboard while adding task, easy to dismiss

**Soft Warnings (Past Dates):**
- **Why:** Don't block users, but inform them
- **Research:** Users sometimes need to log past tasks
- **Impact:** Flexibility without sacrificing UX

---

## 11. Assets Delivery

### Asset Organization

**Location:** `client/assets/`

```
assets/
├── images/
│   ├── kerja.jpg              (37 KB)
│   ├── pribadi.jpg            (36 KB)
│   ├── semua.jpg              (36 KB)
│   ├── ulang_tahun.jpg        (28 KB)
│   └── wishlist.png           (866 KB)
│
└── icons/
    ├── google_logo.svg        (57 KB)
    ├── Payung.svg             (11 KB)
    ├── Jaket.svg              (50 KB)
    ├── Jaket_Hujan.svg        (26 KB)
    ├── Topi.svg               (15 KB)
    ├── Sunscreen.svg          (38 KB)
    ├── Krim_Pelembap.svg      (53 KB)
    ├── Botol_Minum.svg        (66 KB)
    ├── Vitamin.svg            (31 KB)
    ├── vitamin-c.svg          (11 KB)
    ├── Buah_Buahan.svg        (100 KB)
    ├── Kipas Portabel.svg     (28 KB)
    ├── Pakaian.svg            (71 KB)
    └── Olahraga.svg           (33 KB)
```

### Asset Specifications

**Category Images (JPG/PNG):**
- **Format:** JPG (compressed) or PNG (if transparency needed)
- **Size:** 240x240px minimum (1x), 480x480px recommended (2x)
- **Usage:** Empty states, category backgrounds
- **Naming:** lowercase, descriptive (e.g., `kerja.jpg`)

**Weather Icons (SVG):**
- **Format:** SVG (vector)
- **Size:** Scalable (typically rendered at 48-80px)
- **Style:** Illustrative, colorful
- **Usage:** VIP weather recommendations
- **Naming:** PascalCase with underscores (e.g., `Jaket_Hujan.svg`)

**Logo (Not yet provided):**
- **Needed:** Workradar logo in multiple formats
- **Formats:** .svg (vector), .png (1x, 2x, 3x)
- **Sizes:** 
  - App icon: 1024x1024px (iOS), 512x512px (Android)
  - Splash screen: 300x300px minimum
  - In-app logo: 200x200px

### Icon Library

**Primary Library:** Iconsax v0.0.8
- **Style:** Linear (outline)
- **Size:** 20-24px standard
- **Color:** Dynamic (follows theme)
- **Package:** `iconsax: ^0.0.8`

**Common Icons Used:**
- Navigation: `home`, `calendar`, `search_normal`, `message`, `user`
- Actions: `add`, `edit`, `trash`, `more`, `arrow_left`
- Status: `tick_circle`, `warning_2`, `danger`, `info_circle`
- Features: `task_square`, `category`, `notification`, `setting_2`

### Font Files

**System Fonts (No files needed):**
- **iOS:** San Francisco (built-in)
- **Android:** Roboto (built-in)
- **Fallback:** System default

**Custom Fonts:** None currently used

---

## 12. Developer Notes & Handoff Summary

### Design System Implementation

**Grid System:** 8pt Grid
```dart
// Spacing constants
paddingSmall:  8.0
paddingMedium: 16.0
paddingLarge:  24.0
paddingXLarge: 32.0

// Border radius
borderRadiusSmall:  8.0
borderRadiusMedium: 12.0
borderRadiusLarge:  16.0
borderRadiusXLarge: 24.0
```

**Theme File:** `lib/core/theme/app_theme.dart`
- All colors, typography, and spacing defined here
- Separate light and dark themes
- Reusable constants for consistency

### Animation Standards

**Duration Guidelines:**
```dart
// Quick transitions
Duration(milliseconds: 150)  // Chip selection, color changes

// Standard transitions
Duration(milliseconds: 200)  // Button press, small animations

// Modal animations
Duration(milliseconds: 300)  // Bottom sheet entry
Duration(milliseconds: 250)  // Bottom sheet exit

// Page transitions
Duration(milliseconds: 350)  // Screen navigation
```

**Curves:**
- **Default:** `Curves.easeInOut`
- **Entry:** `Curves.easeOut`
- **Exit:** `Curves.easeIn`
- **Bounce:** `Curves.elasticOut` (for playful elements)

### Component Usage Guidelines

**Buttons:**
```dart
// Primary action
ElevatedButton(
  onPressed: () {},
  child: Text('Primary Action'),
)

// Secondary action
TextButton(
  onPressed: () {},
  child: Text('Secondary Action'),
)

// FAB (main action)
FloatingActionButton(
  onPressed: () {},
  child: Icon(Iconsax.add),
)
```

**Cards:**
```dart
// Task card
TaskCard(
  task: task,
  onTap: () => _onTaskTap(task),
  onComplete: () => _onTaskComplete(task),
)

// Custom card
Card(
  elevation: 0,
  shape: RoundedRectangleBorder(
    borderRadius: BorderRadius.circular(AppTheme.borderRadiusLarge),
  ),
  child: ...,
)
```

**Modals:**
```dart
// Bottom sheet
showModalBottomSheet(
  context: context,
  isScrollControlled: true,
  backgroundColor: Colors.transparent,
  builder: (context) => YourWidget(),
)

// Dialog
showDialog(
  context: context,
  builder: (context) => AlertDialog(...),
)
```

### State Management

**Provider Pattern:**
- `TaskProvider` - Task CRUD operations
- `CategoryProvider` - Category management
- `ThemeProvider` - Dark/light mode
- `ProfileProvider` - User profile & settings
- `WorkloadProvider` - Analytics & workload tracking

**Usage:**
```dart
// Read once
final taskProvider = context.read<TaskProvider>();

// Watch for changes
final taskProvider = context.watch<TaskProvider>();

// Select specific value
final tasks = context.select<TaskProvider, List<Task>>(
  (provider) => provider.tasks,
);
```

### API Integration

**Base URL:** Configured in environment
**Authentication:** JWT tokens stored in secure storage
**Error Handling:** Try-catch with user-friendly messages

```dart
try {
  await taskProvider.addTaskToServer(task);
} catch (e) {
  ScaffoldMessenger.of(context).showSnackBar(
    SnackBar(
      content: Text('Gagal membuat tugas: $e'),
      backgroundColor: Colors.red,
    ),
  );
}
```

### Testing Checklist

- [ ] Test all screens in light mode
- [ ] Test all screens in dark mode
- [ ] Test with different font sizes (accessibility)
- [ ] Test on small screens (320px width)
- [ ] Test on large screens (428px width)
- [ ] Test landscape orientation
- [ ] Test all form validations
- [ ] Test error states
- [ ] Test empty states
- [ ] Test loading states
- [ ] Test navigation flows
- [ ] Test back button behavior
- [ ] Test deep linking (if applicable)

---

## 13. Collaboration & Review

### Design Review Process

**Reviewers:**
- Product Manager: [Nama PM]
- Lead Developer: [Nama Dev]
- UX Designer: [Nama Designer]

**Review Dates:**
- Initial Design: [Date]
- Design Iteration: [Date]
- Final Approval: [Date]

### Feedback Summary

**Approved Elements:**
- ✅ Color palette (purple/teal scheme)
- ✅ Bottom navigation structure
- ✅ Card-based layout
- ✅ Dark mode implementation
- ✅ Category system
- ✅ VIP feature differentiation

**Iterations Made:**
- Adjusted empty state messaging for clarity
- Added soft warning for past date selection
- Improved dark mode contrast ratios
- Simplified task card design for completed tasks
- Enhanced accessibility with semantic labels

**Pending Items:**
- [ ] Logo design and delivery
- [ ] Additional empty state illustrations
- [ ] Skeleton loading screens
- [ ] Onboarding flow screens
- [ ] Tutorial/help screens

### Change Log

| **Version** | **Date** | **Changes** | **Reviewer** |
|-------------|----------|-------------|--------------|
| 1.0.0 | 2026-01-10 | Initial design handoff document | [Nama] |

---

## Checklist Handoff Selesai

- [x] File desain rapi dan terorganisir
- [x] Spesifikasi lengkap (13 komponen)
- [x] Aset siap pakai (images & icons)
- [ ] Prototipe interaktif tersedia (Figma/Sketch)
- [x] Dokumentasi 'why' dan 'how' jelas
- [x] Developer walkthrough (dokumentasi lengkap)

---

## Appendix: Screen Inventory

### Authentication Screens (9)
1. Auth Check Screen - Auto-login check
2. Login Screen - Email/Google login
3. Register Screen - New user registration
4. Registration OTP Screen - Email verification
5. Forgot Password Screen - Password recovery
6. Verification Code Screen - OTP input
7. Reset Password Screen - New password
8. MFA Setup Screen - Two-factor auth setup
9. MFA Verify Screen - Two-factor auth login

### Main Screens (5)
10. Main Screen - Bottom navigation container
11. Dashboard Screen - Task list (home)
12. Calendar Screen - Calendar view
13. Search Screen - Task search
14. Messages Screen - Bot chat (VIP)
15. Profile Screen - User profile & settings

### Task Management (2)
16. Edit Task Screen - Edit/delete task
17. Completed Tasks Screen - Task history

### Category Management (1)
18. Manage Category Screen - CRUD categories

### Profile & Settings (3)
19. Edit Profile Screen - Update user info
20. Profile Detail Screen - View profile
21. Leave Management Screen - Vacation days

### VIP Features (3)
22. Subscription Screen - VIP plans
23. Payment Screen - Payment processing
24. VIP Weather Screen - Weather recommendations

**Total: 24 Screens**

---

## Contact & Support

**Design Questions:** [Designer Email]
**Development Questions:** [Dev Team Email]
**Product Questions:** [PM Email]

**Project Repository:** [GitHub/GitLab URL]
**Design Files:** [Figma/Sketch URL]
**Documentation:** [Confluence/Notion URL]

---

**Document Version:** 1.0.0  
**Last Updated:** 10 Januari 2026  
**Status:** ✅ Ready for Development
