package rest

import (
	"fmt"
	"log"

	domain "newnok6/logic-test/pie-fire-dire/internal/core/domain"
	services "newnok6/logic-test/pie-fire-dire/internal/core/services"

	"github.com/gofiber/fiber/v3"
)

type HttpRest struct {
	restServer *fiber.App
	addr       string
	services.ProcessFileService
}

func NewHttpRest(addr string) *HttpRest {

	fileName := "file.txt"
	filePath := "/Users/panupak/Projects/logic-test/pie-fire-dire/files"
	fileType := "txt"

	fileMeta := domain.FileMeta{
		FileName: fileName,
		FilePath: filePath,
		FileType: fileType,
	}

	processFileService := services.NewProcessFileService(fileMeta)

	fmt.Println("File Name:", processFileService.GetFileName())
	fmt.Println("File Path:", processFileService.GetFilePath())

	return &HttpRest{
		restServer:         fiber.New(),
		addr:               addr,
		ProcessFileService: processFileService,
	}

}

func (h *HttpRest) Start() {
	log.Println("Starting HTTP server on port 8080")
	h.registerRoute()
	if err := h.restServer.Listen(h.addr); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
	log.Println("HTTP server started on port 8080")
}

func (h *HttpRest) Stop() {
	log.Println("Stopping HTTP server")
	if err := h.restServer.Shutdown(); err != nil {
		log.Fatalf("Failed to stop HTTP server: %v", err)
	}
	log.Println("HTTP server stopped")
}

func (h *HttpRest) registerRoute() {
	h.restServer.Get("/api/beef/summary", func(c fiber.Ctx) error {
		meatList, err := h.GetMeatList()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get meat list",
			})
		}
		meatResponse := &MeatResponse{
			meatList,
		}

		return c.JSON(meatResponse)
	})
}
