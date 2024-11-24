use std::{collections::HashMap, process::exit};

use shared::{logger::{Logger, LoggerImpl}, token::{token::TokenImpl, token_type::TokenType}};

use crate::shared::{function::{Function, FunctionImpl}, value_name::value_name::VALUE_NAME_REGEX};

pub fn analyze_functions(
    // Pass by reference to avoid moving the value or cloning it 
    functions: &HashMap<String, Function>
) {

    // First check if the main function is even defined
    if !functions.contains_key("main") {
        Logger::err(
            "Main function not defined",
            &[
                "The main function must be defined in the code"
            ],
            &[]
        );

        exit(1);
    }

    // The return type of the main function must be Nothing
    let main_function = functions.get("main").unwrap();
    if main_function.get_return_type() != &TokenType::Nothing {
        Logger::err(
            "Invalid return type for main function",
            &[
                "The return type of the main function must be Nothing"
            ],
            &[
                main_function.get_trace().build_trace().as_str()
            ]
        );

        exit(1);
    }

    for (name, function) in functions.iter() {
        if !VALUE_NAME_REGEX.is_match(name.as_str()).unwrap_or(false) {
            Logger::err(
                format!("Invalid function name: {}", name).as_str(),
                &[
                    "Function names must start with a letter or an underscore",
                ],
                &[
                    function.get_trace().build_trace().as_str()
                ]
            );

            exit(1);
        }

        // The lexer should have already checked if the function name is repeated
        // so we don't need to check it here
    }

}