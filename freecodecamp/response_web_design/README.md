## 📚 **[TABLE OF CONTENTS]()**

### **Responsive Web Design**

> -   ##### In this Responsive Web Design challenge, you'll learn the languages that developers use to build webpages: HTML (Hypertext Markup Language) for content, and CSS (Cascading Style Sheets) for design.
> -   ##### First, you'll build a cat photo app to learn the basics of HTML and CSS. Later, you'll learn modern techniques like CSS variables by building a penguin, and best practices for accessibility by building a web form.
> -   ##### Finally, you'll learn how to make webpages that respond to different screen sizes by building a Twitter card with Flexbox, and a complex blog layout with CSS Grid.

| [**Basic HTML and HTML5**](/HTML%20basic/)                                                                                                                                                                                                                                                                             | [**Basic CSS and CSS3**](/CSS%20basic/)                                                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| &bull; HTML is a markup language that uses a special syntax or notation to describe the structure of a webpage to the browser. HTML elements usually have opening and closing tags that surround and give meaning to content. For example, different elements can describe text as a heading, paragraph, or list item. | &bull; CSS, or Cascading Style Sheets, tell the browser how to display the text and other content that you write in HTML. With CSS, you can control the color, font, size, spacing, and many other aspects of HTML elements. |
| &bull; In this course, you'll build a cat photo app to learn some of the most common HTML elements — the building blocks of any webpage.                                                                                                                                                                               | &bull; Now that you've described the structure of your cat photo app, give it some style with CSS.                                                                                                                           |

| [**Applied Visual Design**](/Applied%20Visual%20Design/)                                                                                           | [**Applied Accessibility**](/Applied%20Accessbility/)                                                                                                                                                                                                 |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| &bull; Visual design is a combination of typography, color theory, graphics, animation, page layout, and more to help deliver your unique message. | &bull; In web development, accessibility refers to web content and a UI (user interface) that can be understood, navigated, and interacted with by a broad audience. This includes people with visual, auditory, mobility, or cognitive disabilities. |
| &bull; In this course, you'll learn how to apply these different elements of visual design to your webpages.                                       | &bull; In this course, you'll learn best practices for building webpages that are accessible to everyone.                                                                                                                                             |

| [**Responsive Web Design Principles**](/Responsive%20Web%20Design%20Principles/)                                                                                                                                                              | [**CSS Flexbox**](/CSS%20Flexbox/)                                                                                                                                                                                                                   |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| &bull; There are many devices that can access the web, and they come in all shapes and sizes. Responsive web design is the practice of designing flexible websites that can respond to different screen sizes, orientations, and resolutions. | &bull; Flexbox is a powerful, well-supported layout method that was introduced with the latest version of CSS, CSS3. With flexbox, it's easy to center elements on the page and create dynamic user interfaces that shrink and expand automatically. |
| &bull; In this course, you'll learn how to use CSS to make your webpages look good, no matter what device they're viewed on.                                                                                                                  | &bull; In this course, you'll learn the fundamentals of flexbox and dynamic layouts by building a Twitter card.                                                                                                                                      |

| [**CSS Grid**](/CSS%20Grid/)                                                                                                                                                                        | [**Responsive Web Design Projects**](/Responsive%20Web%20design%20Projects/)                                                                                                                                                              |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| &bull; The CSS grid is a newer standard that makes it easy to build complex responsive layouts. It works by turning an HTML element into a grid, and lets you place child elements anywhere within. | &bull; Time to put your newly learnt skills to work. By working on these projects, you will get a chance to apply all of the skills, principles, and concepts you have learned so far: HTML, CSS, Visual Design, Accessibility, and more. |
| &bull; In this course, you'll learn the fundamentals of CSS grid by building different complex layouts, including a blog.                                                                           |

## Head

> **Notes:** You can find [a list of everything](https://github.com/joshbuchea/HEAD) that could be found in the `<head>` of an HTML document.

### Meta tag

-   [ ] **Doctype:** The Doctype is HTML5 and is at the top of all your HTML pages.

```html
<!DOCTYPE html>
```

> -   📖 [Determining the character encoding - HTML5 W3C](https://www.w3.org/TR/html5/syntax.html#determining-the-character-encoding)

_The next 2 meta tags (Charset and Viewport) need to come first in the head._

-   [ ] **Charset:** The charset (UTF-8) is declared correctly.

```html
<!-- Set character encoding for the document -->
<meta charset="utf-8" />
```

-   [ ] **Viewport:** The viewport is declared correctly.

```html
<meta
    name="viewport"
    content="width=device-width, initial-scale=1, viewport-fit=cover"
/>
```

-   [ ] **Title:** A title is used on all pages (SEO: Google calculates the pixel width of the characters used in the title, and it cuts off between 472 and 482 pixels. The average character limit would be around 55-characters).

```html
<!-- Document Title -->
<title>Page Title less than 55 characters</title>
```

> -   📖 [Title - HTML - MDN](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title)
> -   🛠 [SERP Snippet Generator](https://www.sistrix.com/serp-snippet-generator/)

-   [ ] **Description:** A meta description is provided, it is unique and doesn't possess more than 150 characters.

```html
<meta
    name="description"
    content="Description of the page less than 150 characters"
/>
```

> -   📖 [Meta Description - HTML - MDN](https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/The_head_metadata_in_HTML#Adding_an_author_and_description)

-   [ ] **Favicons:** Each favicon has been created and displays correctly. If you have only a `favicon.ico`, put it at the root of your site. Normally you won't need to use any markup. However, it's still good practice to link to it using the example below. Today, **PNG format is recommended** over `.ico` format (dimensions: 32x32px).

```html
<!-- Standard favicon -->
<link rel="icon" type="image/x-icon" href="https://example.com/favicon.ico" />
<!-- Recommended favicon format -->
<link rel="icon" type="image/png" href="https://example.com/favicon.png" />
```

> -   🛠 [Favicon Generator](https://www.favicon-generator.org/)
> -   🛠 [RealFaviconGenerator](https://realfavicongenerator.net/)
> -   📖 [Favicon Cheat Sheet](https://github.com/audreyr/favicon-cheat-sheet)
> -   📖 [Favicons, Touch Icons, Tile Icons, etc. Which Do You Need? - CSS Tricks](https://css-tricks.com/favicon-quiz/)
> -   📖 [PNG favicons - caniuse](https://caniuse.com/#feat=link-icon-png)

-   [ ] **Apple Web App Meta:** Apple meta-tags are present.

```html
<!-- Apple Touch Icon (at least 200x200px) -->
<link rel="apple-touch-icon" href="/custom-icon.png" />

<!-- To run the web application in full-screen -->
<meta name="apple-mobile-web-app-capable" content="yes" />

<!-- Has no effect unless you have the previous meta tag -->
<meta name="apple-mobile-web-app-status-bar-style" content="black" />
```

> -   📖 [Configuring Web Applications](https://developer.apple.com/library/content/documentation/AppleApplications/Reference/SafariWebContent/ConfiguringWebApplications/ConfiguringWebApplications.html)
> -   📖 [Supported Meta Tags](https://developer.apple.com/library/content/documentation/AppleApplications/Reference/SafariHTMLRef/Articles/MetaTags.html)

-   [ ] **Windows Tiles:** Windows tiles are present and linked.

```html
<!-- Microsoft Tiles -->
<meta name="msapplication-config" content="browserconfig.xml" />
```

Minimum required xml markup for the `browserconfig.xml` file is as follows:

```xml
<?xml version="1.0" encoding="utf-8"?>
<browserconfig>
   <msapplication>
     <tile>
        <square70x70logo src="small.png"/>
        <square150x150logo src="medium.png"/>
        <wide310x150logo src="wide.png"/>
        <square310x310logo src="large.png"/>
     </tile>
   </msapplication>
</browserconfig>
```

> -   📖 [Browser configuration schema reference](<https://msdn.microsoft.com/en-us/library/dn320426(v=vs.85).aspx>)

-   [ ] **Canonical:** Use `rel="canonical"` to avoid duplicate content.

```html
<!-- Helps prevent duplicate content issues -->
<link
    rel="canonical"
    href="http://example.com/2017/09/a-new-article-to-read.html"
/>
```

> -   📖 [Use canonical URLs - Search Console Help - Google Support](https://support.google.com/webmasters/answer/139066?hl=en)
> -   📖 [5 common mistakes with rel=canonical - Google Webmaster Blog](https://webmasters.googleblog.com/2013/04/5-common-mistakes-with-relcanonical.html)

### HTML tags

-   [ ] **Language attribute:** The `lang` attribute of your website is specified and related to the language of the current page.

```html
<html lang="en"></html>
```

-   [ ] **Direction attribute:** The direction of lecture is specified on the html tag (It can be used on another HTML tag).

```html
<html dir="rtl"></html>
```

> -   📖 [dir - HTML - MDN](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/dir)

-   [ ] **Alternate language:** The language tag of your website is specified and related to the language of the current page.

```html
<link rel="alternate" href="https://es.example.com/" hreflang="es" />
```

-   [ ] **x-default:** The language tag of your website for international landing pages.

```html
<link rel="alternate" href="https://example.com/" hreflang="x-default" />
```

> -   📖 [x-default - Google](https://webmasters.googleblog.com/2013/04/x-default-hreflang-for-international-pages.html)

-   [ ] **Conditional comments:** Conditional comments are present for IE if needed.

> -   📖 [About conditional comments (Internet Explorer) - MSDN - Microsoft](<https://msdn.microsoft.com/en-us/library/ms537512(v=vs.85).aspx>)

-   [ ] **RSS feed:** If your project is a blog or has articles, an RSS link was provided.

-   [ ] **CSS Critical:** The CSS critical (or "above the fold") collects all the CSS used to render the visible portion of the page. It is embedded before your principal CSS call and between `<style></style>` in a single line (minified).

> -   🛠 [Critical by Addy Osmani on GitHub](https://github.com/addyosmani/critical) automates this.

-   [ ] **CSS order:** All CSS files are loaded before any JavaScript files in the `<head>`. (Except the case where sometimes JS files are loaded asynchronously on top of your page).

### Social meta

Visualize and generate automatically our social meta tags with [Meta Tags](https://metatags.io/)

**_Facebook OG_** and **_Twitter Cards_** are, for any website, highly recommended. The other social media tags can be considered if you target a particular presence on those and want to ensure the display.

-   [ ] **Facebook Open Graph:** All Facebook Open Graph (OG) are tested and no one is missing or with false information. Images need to be at least 600 x 315 pixels, although 1200 x 630 pixels is recommended.

> **Notes:** Using `og:image:width` and `og:image:height` will specify the image dimensions to the crawler so that it can render the image immediately without having to asynchronously download and process it.

```html
<meta property="og:type" content="website" />
<meta property="og:url" content="https://example.com/page.html" />
<meta property="og:title" content="Content Title" />
<meta property="og:image" content="https://example.com/image.jpg" />
<meta property="og:description" content="Description Here" />
<meta property="og:site_name" content="Site Name" />
<meta property="og:locale" content="en_US" />
<!-- Next tags are optional but recommended -->
<meta property="og:image:width" content="1200" />
<meta property="og:image:height" content="630" />
```

> -   📖 [A Guide to Sharing for Webmasters](https://developers.facebook.com/docs/sharing/webmasters/)
> -   📖 [Best Practices - Sharing](https://developers.facebook.com/docs/sharing/best-practices/)
> -   🛠 Test your page with the [Facebook OG testing](https://developers.facebook.com/tools/debug/)

-   [ ] **Twitter Card:**

```html
<meta name="twitter:card" content="summary" />
<meta name="twitter:site" content="@site_account" />
<meta name="twitter:creator" content="@individual_account" />
<meta name="twitter:url" content="https://example.com/page.html" />
<meta name="twitter:title" content="Content Title" />
<meta
    name="twitter:description"
    content="Content description less than 200 characters"
/>
<meta name="twitter:image" content="https://example.com/image.jpg" />
```

> -   📖 [Getting started with cards — Twitter Developers](https://developer.twitter.com/en/docs/tweets/optimize-with-cards/guides/getting-started)
> -   🛠 Test your page with the [Twitter card validator](https://cards-dev.twitter.com/validator)

**[⬆ back to top](#-Introduction)**

---

## HTML

### Best practices

-   [ ] **HTML5 Semantic Elements:** HTML5 Semantic Elements are used appropriately (header, section, footer, main...).

> -   📖 [HTML Reference](http://htmlreference.io/)

-   [ ] **Error pages:** Error 404 page and 5xx exist. Remember that the 5xx error pages need to have their CSS integrated (no external call on the current server).

-   [ ] **Noopener:** In case you are using external links with `target="_blank"`, your link should have a `rel="noopener"` attribute to prevent tab nabbing. If you need to support older versions of Firefox, use `rel="noopener noreferrer"`.

> -   📖 [About rel=noopener](https://mathiasbynens.github.io/rel-noopener/)

-   [ ] **Clean up comments:** Unnecessary code needs to be removed before sending the page to production.

### HTML testing

-   [ ] **W3C compliant:** All pages need to be tested with the W3C validator to identify possible issues in the HTML code.

> -   🛠 [W3C validator](https://validator.w3.org/)

-   [ ] **HTML Lint:** I use tools to help me analyze any issues I could have on my HTML code.

> -   🛠 [Dirty markup](https://www.10bestdesign.com/dirtymarkup/)

> -   🛠 [webhint](https://webhint.io/)

-   [ ] **Link checker:** There are no broken links in my page, verify that you don't have any 404 error.

> -   🛠 [W3C Link Checker](https://validator.w3.org/checklink)

-   [ ] **Adblockers test:** Your website shows your content correctly with adblockers enabled (You can provide a message encouraging people to disable their adblocker).

> -   📖 [Use AdBlocking in your Dev Environment](https://andreicioara.com/use-adblocking-in-your-dev-environment-48db500d9b86)

**[⬆ back to top](#-Introduction)**

---

## Webfonts

> **Notes:** Using web fonts may cause Flash Of Unstyled Text/Flash Of Invisible Text - consider having fallback fonts and/or utilizing web font loaders to control behavior.
>
> -   📖 [Google Technical considerations about webfonts](https://developers.google.com/fonts/docs/technical_considerations)

-   [ ] **Webfont format:** WOFF, WOFF2 and TTF are supported by all modern browsers.

> -   📖 [WOFF - Web Open Font Format - Caniuse](https://caniuse.com/#feat=woff).
> -   📖 [WOFF 2.0 - Web Open Font Format - Caniuse](https://caniuse.com/#feat=woff2).
> -   📖 [TTF/OTF - TrueType and OpenType font support](https://caniuse.com/#feat=ttf)
> -   📖 [Using @font-face - CSS-Tricks](https://css-tricks.com/snippets/css/using-font-face/)

-   [ ] **Webfont size:** Webfont sizes don't exceed 2 MB (all variants included).

-   [ ] **Webfont loader:** Control loading behavior with a webfont loader

> -   🛠 [Typekit Web Font Loader](https://github.com/typekit/webfontloader)

**[⬆ back to top](#-Introduction)**

---

## CSS

> **Notes:** Take a look at [CSS guidelines](https://cssguidelin.es/) and [Sass Guidelines](https://sass-guidelin.es/) followed by most Front-End developers. If you have a doubt about CSS properties, you can visit [CSS Reference](http://cssreference.io/). There is also a short [Code Guide](http://codeguide.co/) for consistency.

-   [ ] **Responsive Web Design:** The website is using responsive web design.
-   [ ] **CSS Print:** A print stylesheet is provided and is correct on each page.
-   [ ] **Preprocessors:** Your project is using a CSS preprocessor (e.g [Sass](http://sass-lang.com/), [Less](http://lesscss.org/), [Stylus](http://stylus-lang.com/)).
-   [ ] **Unique ID:** If IDs are used, they are unique to a page.
-   [ ] **Reset CSS:** A CSS reset (reset, normalize or reboot) is used and up to date. _(If you are using a CSS Framework like Bootstrap or Foundation, a Normalize is already included into it.)_

> -   📖 [Reset.css](https://meyerweb.com/eric/tools/css/reset/)
> -   📖 [Normalize.css](https://necolas.github.io/normalize.css/)
> -   📖 [Reboot](https://getbootstrap.com/docs/4.0/content/reboot/)

-   [ ] **JS prefix:** All classes (or id- used in JavaScript files) begin with **js-** and are not styled into the CSS files.

```html
<div id="js-slider" class="my-slider">
    <!-- Or -->
    <div id="id-used-by-cms" class="js-slider my-slider"></div>
</div>
```

-   [ ] **embedded or inline CSS:** Avoid at all cost embedding CSS in `<style>` tags or using inline CSS: only use for valid reasons (e.g. background-image for slider, critical CSS).
-   [ ] **Vendor prefixes:** CSS vendor prefixes are used and are generated accordingly with your browser support compatibility.

> -   🛠 [Autoprefixer CSS online](https://autoprefixer.github.io/)

### Performance

-   [ ] **Concatenation:** CSS files are concatenated in a single file _(Not for HTTP/2)_.
-   [ ] **Minification:** All CSS files are minified.
-   [ ] **Non-blocking:** CSS files need to be non-blocking to prevent the DOM from taking time to load.

> -   📖 [loadCSS by filament group](https://github.com/filamentgroup/loadCSS)
> -   📖 [Example of preload CSS using loadCSS](https://gist.github.com/thedaviddias/c24763b82b9991e53928e66a0bafc9bf)

-   [ ] **Unused CSS:** Remove unused CSS.

> -   🛠 [UnCSS Online](https://uncss-online.com/)
> -   🛠 [PurifyCSS](https://github.com/purifycss/purifycss)
> -   🛠 [PurgeCSS](https://github.com/FullHuman/purgecss)
> -   🛠 [Chrome DevTools Coverage](https://developer.chrome.com/docs/devtools/coverage/)

### CSS testing

-   [ ] **Stylelint:** All CSS or SCSS files are without any errors.

> -   🛠 [stylelint, a CSS linter](https://stylelint.io/)
> -   📖 [Sass guidelines](https://sass-guidelin.es/)

-   [ ] **Responsive web design:** All pages were tested at the following breakpoints: 320px, 768px, 1024px (can be more / different according to your analytics).
        **Responsive Checker -**

    > -   🛠 [Am I Responsive?](http://ami.responsivedesign.is/)
    > -   🛠 [Mobile Friendly Test](https://search.google.com/test/mobile-friendly)
    > -   🛠 [Responsive Website Design Tester](https://responsivedesignchecker.com/)
    > -   🛠 [Responsinator](https://www.responsinator.com/)
    > -   🛠 [XRespond](https://xrespond.com/)

-   [ ] **CSS Validator:** The CSS was tested and pertinent errors were corrected.

> -   🛠 [CSS Validator](https://jigsaw.w3.org/css-validator/)

-   [ ] **Desktop Browsers:** All pages were tested on all current desktop browsers (Safari, Firefox, Chrome, Internet Explorer, EDGE...).
-   [ ] **Mobile Browsers:** All pages were tested on all current mobile browsers (Native browser, Chrome, Safari...).
-   [ ] **OS:** All pages were tested on all current OS (Windows, Android, iOS, Mac...).

-   [ ] **Design fidelity:** Depending on the project and the quality of the creatives, you may be asked to be close to the design. You can use some tools to compare creatives with your code implementation and ensure consistency.

> [Pixel Perfect - Chrome Extension](https://chrome.google.com/webstore/detail/perfectpixel-by-welldonec/dkaagdgjmgdmbnecmcefdhjekcoceebi?hl=en)

-   [ ] **Reading direction:** All pages need to be tested for LTR and RTL languages if they need to be supported.

> -   📖 [Building RTL-Aware Web Apps & Websites: Part 1 - Mozilla Hacks](https://hacks.mozilla.org/2015/09/building-rtl-aware-web-apps-and-websites-part-1/)
> -   📖 [Building RTL-Aware Web Apps & Websites: Part 2 - Mozilla Hacks](https://hacks.mozilla.org/2015/10/building-rtl-aware-web-apps-websites-part-2/)

**[⬆ back to top](#-Introduction)**

---

## Images

> **Notes:** For a complete understanding of image optimization, check the free ebook **[Essential Image Optimization](https://images.guide/)** from Addy Osmani.

### Best practices

-   [ ] **Optimization:** All images are optimized to be rendered in the browser. WebP format could be used for critical pages (like Homepage).

> -   🛠 [Imagemin](https://github.com/imagemin/imagemin)
> -   🛠 Use [ImageOptim](https://imageoptim.com/) to optimise your images for free.
> -   🛠 Use [KeyCDN Image Processing](https://www.keycdn.com/support/image-processing) for image optimization in real time.
> -   🛠 Use [Kraken.io](https://kraken.io/web-interface) awesome alternative for both png and jpg optimization. Up to 1mb per files on free plan.
> -   🛠 [TinyPNG](https://tinypng.com/) optimises png, apng (animated png) and jpg images with very small loss in quality. Free and paid version available.
> -   🛠 [ZorroSVG](http://quasimondo.com/ZorroSVG/) jpg-like compression for transparent images using svg masking.
> -   🛠 [SVGO](https://github.com/svg/svgo) a Nodejs-based tool for optimizing SVG vector graphics files.
> -   🛠 [SVGOMG](https://jakearchibald.github.io/svgomg/) a web-based GUI version of SVGO for optimising your svgs online.

-   [ ] **Picture/Srcset:** You use picture/srcset to provide the most appropriate image for the current viewport of the user.

> -   📖 [How to Build Responsive Images with srcset](https://www.sitepoint.com/how-to-build-responsive-images-with-srcset/)

-   [ ] **Retina:** You provide layout images 2x or 3x, support retina display.
-   [ ] **Sprite:** Small images are in a sprite file (in the case of icons, they can be in an SVG sprite image).
-   [ ] **Width and Height:** Set `width` and `height` attributes on `<img>` if the final rendered image size is known (can be omitted for CSS sizing).
-   [ ] **Alternative text:** All `<img>` have an alternative text which describes the image visually.

> -   📖 [Alt-texts: The Ultimate Guide](https://axesslab.com/alt-texts/)

-   [ ] **Lazy loading:** Images are lazyloaded (A noscript fallback is always provided).
    > -   🛠 [Native lazy loading polyfill](https://github.com/mfranzke/loading-attribute-polyfill/)

**[⬆ back to top](#-Introduction)**

---

## JavaScript

### Best practices

-   [ ] **JavaScript Inline:** You don't have any JavaScript code inline (mixed with your HTML code).
-   [ ] **Concatenation:** JavaScript files are concatenated.
-   [ ] **Minification:** JavaScript files are minified (you can add the `.min` suffix).

> -   📖 [Minify Resources (HTML, CSS, and JavaScript)](https://developers.google.com/speed/docs/insights/MinifyResources)

-   [ ] **JavaScript security:**

> -   📖 [Guidelines for Developing Secure Applications Utilizing JavaScript](https://www.owasp.org/index.php/DOM_based_XSS_Prevention_Cheat_Sheet#Guidelines_for_Developing_Secure_Applications_Utilizing_JavaScript)

-   [ ] **`noscript` tag:** Use `<noscript>` tag in the HTML body if a script type on the page is unsupported or if scripting is currently turned off in the browser. This will be helpful in client-side rendering heavy apps such as React.js, see [examples](https://webdesign.tutsplus.com/tutorials/quick-tip-dont-forget-the-noscript-element--cms-25498).

```html
<noscript> You need to enable JavaScript to run this app. </noscript>
```

-   [ ] **Non-blocking:** JavaScript files are loaded asynchronously using `async` or deferred using `defer` attribute.

> -   📖 [Remove Render-Blocking JavaScript](https://developers.google.com/speed/docs/insights/BlockingJS)

-   [ ] **Optimized and updated JS libraries:** All JavaScript libraries used in your project are necessary (prefer Vanilla Javascript for simple functionalities), updated to their latest version and don't overwhelm your JavaScript with unnecessary methods.

> -   📖 [You may not need jQuery](http://youmightnotneedjquery.com/)
> -   📖 [Vanilla JavaScript for building powerful web applications](https://plainjs.com/)

-   [ ] **Modernizr:** If you need to target some specific features you can use a custom Modernizr to add classes in your `<html>` tag.

> -   🛠 [Customize your Modernizr](https://modernizr.com/download?setclasses)

### JavaScript testing

-   [ ] **ESLint:** No errors are flagged by ESLint (based on your configuration or standards rules).

> -   📖 [ESLint - The pluggable linting utility for JavaScript and JSX](https://eslint.org/)

**[⬆ back to top](#-Introduction)**

---

## Security

### Scan and check your web site

> -   [securityheaders.io](https://securityheaders.io/)
> -   [Observatory by Mozilla](https://observatory.mozilla.org/)

### Best practices

-   [ ] **HTTPS:** HTTPS is used on every page and for all external content (plugins, images...).

> -   🛠 [Let's Encrypt - Free SSL/TLS Certificates](https://letsencrypt.org/)
> -   🛠 [Free SSL Server Test](https://www.ssllabs.com/ssltest/index.html)
> -   📖 [Strict Transport Security](http://caniuse.com/#feat=stricttransportsecurity)

-   [ ] **HTTP Strict Transport Security (HSTS):** The HTTP header is set to 'Strict-Transport-Security'.

> -   🛠 [Check HSTS preload status and eligibility](https://hstspreload.org/)
> -   📖 [HTTP Strict Transport Security Cheat Sheet - OWASP](https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Strict_Transport_Security_Cheat_Sheet.html)
> -   📖 [Transport Layer Protection Cheat Sheet - OWASP](https://cheatsheetseries.owasp.org/cheatsheets/Transport_Layer_Protection_Cheat_Sheet.html)

-   [ ] **Cross Site Request Forgery (CSRF):** You ensure that requests made to your server-side are legitimate and originate from your website / app to prevent CSRF attacks.

> -   📖 [Cross-Site Request Forgery (CSRF) Prevention Cheat Sheet - OWASP](https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html)

-   [ ] **Cross Site Scripting (XSS):** Your page or website is free from XSS possible issues.

> -   📖 [XSS (Cross Site Scripting) Prevention Cheat Sheet - OWASP](https://cheatsheetseries.owasp.org/cheatsheets/Cross_Site_Scripting_Prevention_Cheat_Sheet.html)
> -   📖 [DOM based XSS Prevention Cheat Sheet - OWASP](https://cheatsheetseries.owasp.org/cheatsheets/DOM_based_XSS_Prevention_Cheat_Sheet.html)

-   [ ] **Content Type Options:** Prevents Google Chrome and Internet Explorer from trying to mime-sniff the content-type of a response away from the one being declared by the server.

> -   📖 [X-Content-Type-Options - Scott Helme](https://scotthelme.co.uk/hardening-your-http-response-headers/#x-content-type-options)

-   [ ] **X-Frame-Options (XFO):** Protects your visitors against clickjacking attacks.

> -   📖 [X-Frame-Options - Scott Helme](https://scotthelme.co.uk/hardening-your-http-response-headers/#x-frame-options)
> -   📖 [RFC7034 - HTTP Header Field X-Frame-Options](https://tools.ietf.org/html/rfc7034)

-   [ ] **Content Security Policy:** Defines how content is loaded on your site and from where it is permitted to be loaded. Can also be used to protect against clickjacking attacks.

> -   📖 [Content Security Policy - An Introduction - Scott Helme](https://scotthelme.co.uk/content-security-policy-an-introduction/)
> -   📖 [CSP Cheat Sheet - Scott Helme](https://scotthelme.co.uk/csp-cheat-sheet/)
> -   📖 [CSP Cheat Sheet - OWASP](https://cheatsheetseries.owasp.org/cheatsheets/Content_Security_Policy_Cheat_Sheet.html)
> -   📖 [Content Security Policy Reference](https://content-security-policy.com/)

**[⬆ back to top](#-Introduction)**

---

## Performance

### Best practices

-   [ ] **Goals to achieve:** Your pages should reach these goals:
    -   First Meaningful Paint under 1 second
    -   Time To Interactive under 5 seconds for the "average" configuration (a $200 Android on a slow 3G network with 400ms RTT and 400kbps transfer speed) and under 2 seconds for repeat visits
    -   Critical file size under 170Kb gzipped

> -   🛠 [Website Page Analysis](https://tools.pingdom.com)
> -   🛠 [WebPageTest](https://www.webpagetest.org/)
> -   📖 [Size Limit: Make the Web lighter](https://evilmartians.com/chronicles/size-limit-make-the-web-lighter)

-   [ ] **Minified HTML:** Your HTML is minified.

-   [ ] **Lazy loading:** Images, scripts and CSS need to be lazy loaded to improve the response time of the current page (See details in their respective sections).

-   [ ] **Cookie size:** If you are using cookies be sure each cookie doesn't exceed 4096 bytes and your domain name doesn't have more than 20 cookies.

> -   📖 [Cookie specification: RFC 6265](https://tools.ietf.org/html/rfc6265)
> -   📖 [Cookies](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies)
> -   🛠 [Browser Cookie Limits](http://browsercookielimits.squawky.net/)

-   [ ] **Third party components:** Third party iframes or components relying on external JS (like sharing buttons) are replaced by static components when possible, thus limiting calls to external APIs and keeping your user's activity private.

> -   🛠 [Simple sharing buttons generator](https://simplesharingbuttons.com/)

### Preparing upcoming requests

> -   📖 [Explanation of the following techniques](https://css-tricks.com/prefetching-preloading-prebrowsing/)

-   [ ] **DNS resolution:** DNS of third-party services that may be needed are resolved in advance during idle time using `dns-prefetch`.

```html
<link rel="dns-prefetch" href="https://example.com" />
```

-   [ ] **Preconnection:** DNS lookup, TCP handshake and TLS negotiation with services that will be needed soon is done in advance during idle time using `preconnect`.

```html
<link rel="preconnect" href="https://example.com" />
```

-   [ ] **Prefetching:** Resources that will be needed soon (e.g. lazy loaded images) are requested in advance during idle time using `prefetch`.

```html
<link rel="prefetch" href="image.png" />
```

-   [ ] **Preloading:** Resources needed in the current page (e.g. scripts placed at the end of `<body>`) in advance using `preload`.

```html
<link rel="preload" href="app.js" />
```

> -   📖 [Difference between prefetch and preload](https://medium.com/reloading/preload-prefetch-and-priorities-in-chrome-776165961bbf)

### Performance testing

-   [ ] **Google PageSpeed:** All your pages were tested (not only the homepage) and have a score of at least 90/100.

> -   🛠 [Google PageSpeed](https://developers.google.com/speed/pagespeed/insights/)
> -   🛠 [Test your mobile speed with Google](https://testmysite.withgoogle.com)
> -   🛠 [WebPagetest - Website Performance and Optimization Test](https://www.webpagetest.org/)
> -   🛠 [GTmetrix - Website speed and performance optimization](https://gtmetrix.com/)
> -   🛠 [Speedrank - Improve the performance of your website](https://speedrank.app/)

**[⬆ back to top](#-Introduction)**

---

## Accessibility

> **Notes:** You can watch the playlist [A11ycasts with Rob Dodson](https://www.youtube.com/playlist?list=PLNYkxOF6rcICWx0C9LVWWVqvHlYJyqw7g) 📹

### Best practices

-   [ ] **Progressive enhancement:** Major functionality like main navigation and search should work without JavaScript enabled.

> -   📖 [Enable / Disable JavaScript in Chrome Developer Tools](https://www.youtube.com/watch?v=kBmvq2cE0D8)

-   [ ] **Color contrast:** Color contrast should at least pass WCAG AA (AAA for mobile).

> -   🛠 [Contrast ratio](https://leaverou.github.io/contrast-ratio/)

#### Headings

-   [ ] **H1:** All pages have an H1 which is not the title of the website.
-   [ ] **Headings:** Headings should be used properly and in the right order (H1 to H6).

> -   📹 [Why headings and landmarks are so important -- A11ycasts #18](https://www.youtube.com/watch?v=vAAzdi1xuUY&index=9&list=PLNYkxOF6rcICWx0C9LVWWVqvHlYJyqw7g)

### Semantics

-   [ ] **Specific HTML5 input types are used:** This is especially important for mobile devices that show customized keypads and widgets for different types.

> -   📖 [Mobile Input Types](http://mobileinputtypes.com/)

### Form

-   [ ] **Label:** A label is associated with each input form element. In case a label can't be displayed, use `aria-label` instead.

> -   📖 [Using the aria-label attribute - MDN](https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-label)

### Accessibility testing

-   [ ] **Accessibility standards testing:** Use the WAVE tool to test if your page respects the accessibility standards.

> -   🛠 [Wave testing](http://wave.webaim.org/)

-   [ ] **Keyboard navigation:** Test your website using only your keyboard in a previsible order. All interactive elements are reachable and usable.
-   [ ] **Screen-reader:** All pages were tested in a screen-reader (VoiceOver, ChromeVox, NVDA or Lynx).
-   [ ] **Focus style:** If the focus is disabled, it is replaced by visible state in CSS.

> -   📹 [Managing Focus - A11ycasts #22](https://www.youtube.com/watch?v=srLRSQg6Jgg&index=5&list=PLNYkxOF6rcICWx0C9LVWWVqvHlYJyqw7g)

**[⬆ back to top](#-Introduction)**

---

## SEO

-   [ ] **Google Analytics:** Google Analytics is installed and correctly configured.

> -   🛠 [Google Analytics](https://analytics.google.com/analytics/web/)
> -   🛠 [GA Checker (and others)](http://www.gachecker.com/)

-   [ ] **Search Console:** Search Console is installed and correctly configured. It is a free service offered by Google that helps you monitor, maintain, and troubleshoot your site's presence in Google Search results.

> -   🛠 [Search Console](https://search.google.com/search-console/about)

-   [ ] **Headings logic:** Heading text helps to understand the content in the current page.

> -   🛠 [Tota11y, tab Headings](http://khan.github.io/tota11y/#Try-it)

-   [ ] **sitemap.xml:** A sitemap.xml exists and was submitted to Google Search Console (previously Google Webmaster Tools).

> -   🛠 [Sitemap generator](https://websiteseochecker.com/html-sitemap-generator/)

-   [ ] **robots.txt:** The robots.txt is not blocking webpages.

> -   📖 [The robots.txt file](https://varvy.com/robottxt.html)
> -   🛠 Test your robots.txt with [Google Robots Testing Tool](https://www.google.com/webmasters/tools/robots-testing-tool)

-   [ ] **Structured Data:** Pages using structured data are tested and are without errors. Structured data helps crawlers understand the content in the current page.

> -   📖 [Introduction to Structured Data - Search - Google Developers](https://developers.google.com/search/docs/guides/intro-structured-data)
> -   📖 [JSON-LD](https://json-ld.org/)
> -   📖 [Microdata](https://www.w3.org/TR/microdata/)
> -   🛠 Test your page with the [Rich Results Test](https://search.google.com/test/rich-results)
> -   🛠 Complete list of vocabularies that can be used as structured data. [Schema.org Full Hierarchy](http://schema.org/docs/full.html)

-   [ ] **Sitemap HTML:** An HTML sitemap is provided and is accessible via a link in the footer of your website.

> -   📖 [Sitemap guidelines - Google Support](https://support.google.com/webmasters/answer/183668?hl=en)

**[⬆ back to top](#-Introduction)**
