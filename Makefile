build_and_run:
	@go build
	@./wapp_start_page_links
all:
	@go build
generate_favicons:
	@/usr/bin/convert -resize 16x16 assets/favicon/favicon.png assets/favicon/favicon-16x16.png
	@/usr/bin/convert -resize 32x32 assets/favicon/favicon.png assets/favicon/favicon-32x32.png
	@/usr/bin/convert -resize x180 assets/favicon/favicon.png assets/favicon/apple-touch-icon-180x180.png
	@/usr/bin/convert -resize x152 assets/favicon/favicon.png assets/favicon/apple-touch-icon-152x152.png
	@/usr/bin/convert -resize x120 assets/favicon/favicon.png assets/favicon/apple-touch-icon-120x120.png
	@/usr/bin/convert -resize x76  assets/favicon/favicon.png assets/favicon/apple-touch-icon-76x76.png
	@/usr/bin/convert -resize x60  assets/favicon/favicon.png assets/favicon/apple-touch-icon-60x60.png