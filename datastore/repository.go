package datastore

type ExperienceSQLRepository interface {
	SaveWarehouse()
}

type ExperienceDocumentRepository interface {
	IndexDocument()
}

type ExperienceRepository interface {
	ExperienceDocumentRepository
	ExperienceSQLRepository
}

type SqlRepository struct {
	SQLGateway
}

type DocumentRepository struct {
	DC DocIndexClient
}

func (sqlRepo *SqlRepository) SaveWarehouse() {
	sqlRepo.Save()
}

func (docClient *DocumentRepository) IndexDocument() {
	docClient.DC.Index()
}
