package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

type Operacao interface {
	Sacar(valor float64) (bool, string)
}

func PagarBoleto(operacao Operacao, valor float64) (bool, string) {
	return operacao.Sacar(valor)
}

func main() {

	conta1 := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "José da Silva",
			CPF:       "224.241.225-87",
			Profissao: "Desenvolvedor"},
		NumeroAgencia: 123,
		NumeroConta:   1234,
	}

	conta2 := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Juliano",
			CPF:       "587.698.655-99",
			Profissao: "RH"},
		NumeroAgencia: 123,
		NumeroConta:   5678,
	}

	fmt.Println(conta1)
	fmt.Println("Realizando saque")
	_, mensagem := conta1.Sacar(150)
	fmt.Println(mensagem)
	fmt.Println("Obtendo o saldo inicial", conta1.GetSaldo())
	valorDeposito := 1000.0
	fmt.Println("Depositando:", valorDeposito)
	_, mensagem = conta1.Depositar(valorDeposito)
	fmt.Println(mensagem)
	fmt.Println("Obtendo o saldo após o depósito", conta1.GetSaldo())

	_, mensagem = conta1.Sacar(150)
	fmt.Println(mensagem)
	fmt.Println("Obtendo o saldo após o saque", conta1.GetSaldo())

	_, mensagem = conta1.Transferir(&conta2, 250)
	fmt.Println(mensagem)
	fmt.Println("Obtendo o saldo após receber uma transferência", conta2.GetSaldo())

	_, mensagem = PagarBoleto(&conta2, 250)
	fmt.Print(mensagem)

}
