package cli

import (
	"github.com/mwinyimoha/commons/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *CMD) AddTodo() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-todo",
		Short:   "Add a new todo item",
		Aliases: []string{"add"},
		RunE: func(cmd *cobra.Command, args []string) error {
			description, err := cmd.Flags().GetString("description")
			if err != nil {
				return errors.WrapError(err, errors.Internal, "failed to get arg")
			}

			todo, err := c.service.CreateTodo(description)
			if err != nil {
				return err
			}

			cmd.Println("Todo added with ID:", todo.ID)
			return nil
		},
	}

	cmd.Flags().String("description", "", "Description of the task to be done")
	return cmd
}

func (c *CMD) ListTodos() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list-todos",
		Short:   "List all todo items",
		Aliases: []string{"list"},
		RunE: func(cmd *cobra.Command, args []string) error {
			todos, err := c.service.GetTodos()
			if err != nil {
				return err
			}

			if len(todos) == 0 {
				cmd.Println("No todos found.")
				return nil
			}

			for _, todo := range todos {
				cmd.Printf("ID: %s, Description: %s, Created: %s\n", todo.ID, todo.Description, todo.CreatedAt)
			}
			return nil
		},
	}

	return cmd
}

func (c *CMD) GetTodo() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get-todo",
		Short:   "Get a todo item by ID",
		Aliases: []string{"get"},
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cmd.Flags().GetString("id")
			if err != nil {
				return errors.WrapError(err, errors.Internal, "failed to get arg")
			}

			todo, err := c.service.GetTodo(id)
			if err != nil {
				return err
			}

			cmd.Printf("ID: %s, Description: %s, Created: %s\n", todo.ID, todo.Description, todo.CreatedAt)
			return nil
		},
	}

	cmd.Flags().String("id", "", "ID of the todo item")
	return cmd
}

func (c *CMD) UpdateTodo() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-todo",
		Short:   "Update a todo item by ID",
		Aliases: []string{"update"},
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cmd.Flags().GetString("id")
			if err != nil {
				return errors.WrapError(err, errors.Internal, "failed to get arg")
			}

			description, err := cmd.Flags().GetString("description")
			if err != nil {
				return errors.WrapError(err, errors.Internal, "failed to get arg")
			}

			todo, err := c.service.UpdateTodo(id, description)
			if err != nil {
				return err
			}

			cmd.Printf("Updated Todo - ID: %s, Description: %s, Created: %s\n", todo.ID, todo.Description, todo.CreatedAt)
			return nil
		},
	}

	cmd.Flags().String("id", "", "ID of the todo item")
	cmd.Flags().String("description", "", "New description of the task to be done")
	return cmd
}

func (c *CMD) DeleteTodo() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete-todo",
		Short:   "Delete a todo item by ID",
		Aliases: []string{"delete"},
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cmd.Flags().GetString("id")
			if err != nil {
				return errors.WrapError(err, errors.Internal, "failed to get arg")
			}

			if err = c.service.DeleteTodo(id); err != nil {
				return err
			}

			cmd.Println("Todo deleted with ID:", id)
			return nil
		},
	}

	cmd.Flags().String("id", "", "ID of the todo item")
	return cmd
}
