package user_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/HeronWest/nostrataskapi/internal/user"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, user.Repository) {
	// Cria o mock do banco de dados
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("Erro ao criar mock de banco de dados")
	}

	// Configura o GORM para usar o mock do banco de dados
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Erro ao criar GORM com mock")
	}

	// Cria o repositório de usuários com o mock do DB
	userRepo := user.NewUserRepository(gormDB)

	// Retorna o GORM DB, o mock e o repositório
	return gormDB, mock, userRepo
}

func TestUserRepositoryWithMock(t *testing.T) {
	// Configura o mock do DB e o repositório
	_, mock, userRepo := setupMockDB()

	// Testando a criação de um usuário
	t.Run("Create User", func(t *testing.T) {
		// Definindo o comportamento esperado do mock para a consulta de criação
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO \"users\"").
			WithArgs("John Doe", "john@example.com", "admin").
			WillReturnResult(sqlmock.NewResult(1, 1)) // Simula a inserção bem-sucedida
		mock.ExpectCommit()

		// Criação do usuário
		newUser := &user.User{Name: "John Doe", Email: "john@example.com", Role: "admin"}
		err := userRepo.Create(newUser)

		// Verificando se não houve erro
		assert.NoError(t, err)

		// Verificando se as expectativas do mock foram atendidas
		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	// Testando a busca de um usuário pelo ID
	t.Run("Find User By ID", func(t *testing.T) {
		// Definindo o comportamento esperado do mock para a consulta de busca por ID
		mock.ExpectQuery("SELECT * FROM \"users\" WHERE \"users\".\"id\" = ?").
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role"}).
				AddRow(1, "John Doe", "john@example.com", "admin"))

		// Buscando o usuário pelo ID
		user, err := userRepo.FindByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "John Doe", user.Name)

		// Verificando se as expectativas do mock foram atendidas
		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	// Testando a atualização de um usuário
	t.Run("Update User", func(t *testing.T) {
		// Definindo o comportamento esperado do mock para a consulta de atualização
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE \"users\" SET \"name\" = ?, \"email\" = ?, \"role\" = ? WHERE \"id\" = ?").
			WithArgs("Jane Doe", "jane@example.com", "user", 1).
			WillReturnResult(sqlmock.NewResult(1, 1)) // Simula a atualização bem-sucedida
		mock.ExpectCommit()

		// Buscando o usuário e atualizando
		userToUpdate := &user.User{ID: 1, Name: "Jane Doe", Email: "jane@example.com", Role: "user"}
		err := userRepo.Update(userToUpdate)
		assert.NoError(t, err)

		// Verificando se as expectativas do mock foram atendidas
		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	// Testando a exclusão de um usuário
	t.Run("Delete User", func(t *testing.T) {
		// Definindo o comportamento esperado do mock para a consulta de deleção
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM \"users\" WHERE \"id\" = ?").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1)) // Simula a exclusão bem-sucedida
		mock.ExpectCommit()

		// Deletando o usuário
		err := userRepo.Delete(1)
		assert.NoError(t, err)

		// Verificando se as expectativas do mock foram atendidas
		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
