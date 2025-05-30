// Copyright 2023 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

// Code generated by "stringer"; DO NOT EDIT.

package schemachange

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[insertRow-0]
	_ = x[selectStmt-1]
	_ = x[validate-2]
	_ = x[renameIndex-3]
	_ = x[renameSequence-4]
	_ = x[renameTable-5]
	_ = x[renameView-6]
	_ = x[alterDatabaseAddRegion-7]
	_ = x[alterDatabasePrimaryRegion-8]
	_ = x[alterDatabaseSurvivalGoal-9]
	_ = x[alterDatabaseAddSuperRegion-10]
	_ = x[alterDatabaseDropSuperRegion-11]
	_ = x[alterFunctionRename-12]
	_ = x[alterFunctionSetSchema-13]
	_ = x[alterTableAddColumn-14]
	_ = x[alterTableAddConstraint-15]
	_ = x[alterTableAddConstraintForeignKey-16]
	_ = x[alterTableAddConstraintUnique-17]
	_ = x[alterTableAlterColumnType-18]
	_ = x[alterTableAlterPrimaryKey-19]
	_ = x[alterTableDropColumn-20]
	_ = x[alterTableDropColumnDefault-21]
	_ = x[alterTableDropConstraint-22]
	_ = x[alterTableDropNotNull-23]
	_ = x[alterTableDropStored-24]
	_ = x[alterTableRLS-25]
	_ = x[alterTableLocality-26]
	_ = x[alterTableRenameColumn-27]
	_ = x[alterTableSetColumnDefault-28]
	_ = x[alterTableSetColumnNotNull-29]
	_ = x[alterTypeDropValue-30]
	_ = x[createTypeEnum-31]
	_ = x[createTypeComposite-32]
	_ = x[createIndex-33]
	_ = x[createSchema-34]
	_ = x[createSequence-35]
	_ = x[createTable-36]
	_ = x[createTableAs-37]
	_ = x[createView-38]
	_ = x[createFunction-39]
	_ = x[createPolicy-40]
	_ = x[commentOn-41]
	_ = x[dropFunction-42]
	_ = x[dropIndex-43]
	_ = x[dropSchema-44]
	_ = x[dropSequence-45]
	_ = x[dropTable-46]
	_ = x[dropView-47]
	_ = x[dropPolicy-48]
}

func (i opType) String() string {
	switch i {
	case insertRow:
		return "insertRow"
	case selectStmt:
		return "selectStmt"
	case validate:
		return "validate"
	case renameIndex:
		return "renameIndex"
	case renameSequence:
		return "renameSequence"
	case renameTable:
		return "renameTable"
	case renameView:
		return "renameView"
	case alterDatabaseAddRegion:
		return "alterDatabaseAddRegion"
	case alterDatabasePrimaryRegion:
		return "alterDatabasePrimaryRegion"
	case alterDatabaseSurvivalGoal:
		return "alterDatabaseSurvivalGoal"
	case alterDatabaseAddSuperRegion:
		return "alterDatabaseAddSuperRegion"
	case alterDatabaseDropSuperRegion:
		return "alterDatabaseDropSuperRegion"
	case alterFunctionRename:
		return "alterFunctionRename"
	case alterFunctionSetSchema:
		return "alterFunctionSetSchema"
	case alterTableAddColumn:
		return "alterTableAddColumn"
	case alterTableAddConstraint:
		return "alterTableAddConstraint"
	case alterTableAddConstraintForeignKey:
		return "alterTableAddConstraintForeignKey"
	case alterTableAddConstraintUnique:
		return "alterTableAddConstraintUnique"
	case alterTableAlterColumnType:
		return "alterTableAlterColumnType"
	case alterTableAlterPrimaryKey:
		return "alterTableAlterPrimaryKey"
	case alterTableDropColumn:
		return "alterTableDropColumn"
	case alterTableDropColumnDefault:
		return "alterTableDropColumnDefault"
	case alterTableDropConstraint:
		return "alterTableDropConstraint"
	case alterTableDropNotNull:
		return "alterTableDropNotNull"
	case alterTableDropStored:
		return "alterTableDropStored"
	case alterTableRLS:
		return "alterTableRLS"
	case alterTableLocality:
		return "alterTableLocality"
	case alterTableRenameColumn:
		return "alterTableRenameColumn"
	case alterTableSetColumnDefault:
		return "alterTableSetColumnDefault"
	case alterTableSetColumnNotNull:
		return "alterTableSetColumnNotNull"
	case alterTypeDropValue:
		return "alterTypeDropValue"
	case createTypeEnum:
		return "createTypeEnum"
	case createTypeComposite:
		return "createTypeComposite"
	case createIndex:
		return "createIndex"
	case createSchema:
		return "createSchema"
	case createSequence:
		return "createSequence"
	case createTable:
		return "createTable"
	case createTableAs:
		return "createTableAs"
	case createView:
		return "createView"
	case createFunction:
		return "createFunction"
	case createPolicy:
		return "createPolicy"
	case commentOn:
		return "commentOn"
	case dropFunction:
		return "dropFunction"
	case dropIndex:
		return "dropIndex"
	case dropSchema:
		return "dropSchema"
	case dropSequence:
		return "dropSequence"
	case dropTable:
		return "dropTable"
	case dropView:
		return "dropView"
	case dropPolicy:
		return "dropPolicy"
	default:
		return "opType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
